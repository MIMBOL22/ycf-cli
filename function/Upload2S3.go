package function

import (
	"crypto/sha256"
	"encoding/hex"
	ysdkFunc "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
	"io"
	"log"
	"os"
	"ycf-cli/config"
	"ycf-cli/internal"
)

func (f *Function) Upload2S3() {
	log.Println("Загрузка в S3 функции", f.PathName)

	localFuncFilename := "./" + config.CACHE_DIR + "/cache/" + f.PathName + "/func.zip"
	destFuncFilename := "ycf-cli/" + f.Config.Id + ".zip"

	file, err := os.Open(localFuncFilename)
	if err != nil {
		log.Println("Ошибка чтения архива функции", f.PathName, "для загрузки в S3:")
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = internal.Upload2S3(config.Config.S3Bucket, destFuncFilename, file)
	if err != nil {
		log.Println("Ошибка загрузки архива функции", f.PathName, ":")
		log.Println(err)
		os.Exit(1)
	}

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Println("Ошибка получения хэша архива функции", f.PathName, ":")
		log.Println(err)
		os.Exit(1)
	}

	f.S3Package = &ysdkFunc.Package{
		BucketName: config.Config.S3Bucket,
		ObjectName: destFuncFilename,
		Sha256:     hex.EncodeToString(h.Sum(nil)),
	}
}
