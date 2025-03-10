package function

import (
	"github.com/evanw/esbuild/pkg/api"
	"log"
	"os"
)

func (f *Function) Build() {
	f.Compile()
	log.Println("Сборка бандла функции", f.PathName)

	entryFile, _, _, outDir := f.getPaths()
	err := os.RemoveAll(outDir)
	if err != nil {
		log.Println("Ошибка очистки папки:", outDir)
		os.Exit(1)
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
		log.Println("Ошибка сборки бандла функции:", result.Errors)
		os.Exit(1)
	}

	f.CopyAdditional()
}
