package api

import (
	"context"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"log"
	"os"
)

type YaClAPI struct {
	Sdk *ycsdk.SDK
	Ctx context.Context
}

func (y *YaClAPI) Init(credentials ycsdk.Credentials) {
	y.Ctx = context.Background()
	sdkInstance, err := ycsdk.Build(y.Ctx, ycsdk.Config{
		Credentials: credentials,
	})
	if err != nil {
		log.Println("Ошибка инициализации YC API:")
		log.Println(err)
		os.Exit(1)
	}

	y.Sdk = sdkInstance
	log.Println("Успешная инициализация Yandex Cloud API")
}
