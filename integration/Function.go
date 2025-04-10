package integration

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"ysd-cli/utils"
)

type Function struct {
	Config      FunctionConfig
	PathName    string
	ApiData     *functions.Function
	S3Package   *functions.Package
	Integration *Integration
}

type FunctionConfigEnvironment struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Value  string `yaml:"value"`
	Secret struct {
		Id        string `yaml:"id"`
		VersionId string `yaml:"version_id"`
		Key       string `yaml:"key"`
	} `yaml:"secret"`
}

type FunctionConfigEntryPoint struct {
	File     string `yaml:"file"`
	Function string `yaml:"function"`
}

type FunctionConfig struct {
	EnvironmentName    string                      `yaml:"name"`
	Id                 string                      `yaml:"id"`
	Memory             int64                       `yaml:"memory"`
	Runtime            string                      `yaml:"runtime"`
	Description        string                      `yaml:"description"`
	Entrypoint         FunctionConfigEntryPoint    `yaml:"entrypoint"`
	EntrypointFunction string                      `yaml:"entrypoint_function"`
	Timeout            int64                       `yaml:"timeout"`
	ServiceAccountId   string                      `yaml:"service_account_id"`
	Environment        []FunctionConfigEnvironment `yaml:"environment"`
	AdditionalFiles    []string                    `yaml:"additional_files"`
}

type FunctionEnvironments struct {
	Environments []FunctionConfig `yaml:"environments"`
}

func (i *Integration) NewFunctionList(funcBaseDir, environment string) ([]Function, error) {
	dirs, err := utils.ListOfDir(funcBaseDir, ConfigFile)
	if err != nil {
		return nil, err
	}

	funcConfigArray := make([]Function, len(dirs))

	for i, funcDir := range dirs {
		yamlFile, err := os.ReadFile(funcBaseDir + "/" + funcDir + "/" + ConfigFile)
		if err != nil {
			return nil, err
		}

		var functionEnvironments FunctionEnvironments
		err = yaml.Unmarshal(yamlFile, &functionEnvironments)
		if err != nil {
			return nil, err
		}

		for _, fEnvironment := range functionEnvironments.Environments {
			if fEnvironment.EnvironmentName == environment {
				funcConfigArray[i].Config = fEnvironment
				break
			}
		}
		funcConfigArray[i].PathName = funcDir
	}

	return funcConfigArray, nil
}

// BuildCode Сборка кода. Не путать со сборкой функции, как интеграции
func (f *Function) BuildCode() error {
	err := f.Compile()
	if err != nil {
		return err
	}
	log.Println("Сборка бандла функции", f.PathName)

	entryFile, _, _, outDir := f.getPaths()
	err = os.RemoveAll(outDir)
	if err != nil {
		return err
	}

	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{entryFile},
		Bundle:            true,
		Write:             true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Platform:          api.PlatformNode,
		Packages:          api.PackagesBundle,
		Format:            api.FormatCommonJS,
		LegalComments:     api.LegalCommentsNone,
		Outdir:            outDir,
	})

	if len(result.Errors) != 0 {
		return errors.New(result.Errors[0].Text)
	}

	err = f.CopyAdditional()
	return err
}

func (f *Function) Compile() error {
	_, entryDir, funcCacheDir, _ := f.getPaths()

	log.Println("TSC сборка функции", f.PathName)

	cmd := exec.Command("tsc", "--outDir", "../../"+funcCacheDir+"/tsc", "-p", "./tsconfig.json")
	cmd.Dir = entryDir
	out, err := cmd.Output()

	if err != nil {
		return errors.New(string(out))
	}
	return nil
}

func (f *Function) CopyAdditional() error {
	_, entryDir, funcCacheDir, _ := f.getPaths()

	log.Println("Копирование дополнительных файлов фунации", f.PathName)

	wgErrors := make(chan error, len(f.Config.AdditionalFiles))

	var wg sync.WaitGroup
	for _, file := range f.Config.AdditionalFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			filePath := funcCacheDir + "/dist/" + file

			// Создание структуры папок для копируемого файла
			parts := strings.Split(filePath, "/")
			onlyDirParts := parts[0 : len(parts)-1]
			onlyDir := strings.Join(onlyDirParts, "/")
			if _, err := os.Stat(onlyDir); errors.Is(err, os.ErrNotExist) {
				err := os.MkdirAll(onlyDir, 0775)
				if err != nil {
					wgErrors <- err
					return
				}
			}

			_, err := utils.CopyFile(entryDir+"/"+file, filePath)
			if err != nil {
				wgErrors <- err
				return
			}
		}(file)
	}

	wg.Wait()

	select {
	case err := <-wgErrors:
		return err
	default:
		wg.Done()
		return nil
	}
}

func (f *Function) Upload2S3(bucket string) error {
	log.Println("Загрузка в S3 функции", f.PathName)

	localFuncFilename := "./" + CacheDir + "/cache/" + f.PathName + "/func.zip"
	destFuncFilename := "ysd-cli/" + f.Config.Id + ".zip"

	file, err := os.Open(localFuncFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = utils.Upload2S3(bucket, destFuncFilename, file)
	if err != nil {
		return err
	}

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return err
	}

	f.S3Package = &functions.Package{
		BucketName: bucket,
		ObjectName: destFuncFilename,
		Sha256:     hex.EncodeToString(h.Sum(nil)),
	}
	return nil
}

func (f *Function) Zip() error {
	_, _, funcCacheDir, _ := f.getPaths()

	err := utils.Zip(funcCacheDir+"/dist", funcCacheDir+"/func.zip")
	return err
}

func (f *Function) getPaths() (string, string, string, string) {
	entryDir := f.Integration.Dir + "/" + f.PathName
	entryFile := entryDir + "/" + f.Config.Entrypoint.File
	funcCacheDir := "./" + CacheDir + "/cache/functions/" + f.PathName
	outDir := funcCacheDir + "/dist"
	return entryFile, entryDir, funcCacheDir, outDir
}
