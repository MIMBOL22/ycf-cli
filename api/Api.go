package api

import (
	"context"
	"errors"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"log"
	"os"
)

type YaClAPI struct {
	Sdk *ycsdk.SDK
	Ctx context.Context
}

type AuthData struct {
	ServiceAccountKeyPath string `yaml:"service_account_key_path"`
	Type                  string `yaml:"auth_type"`
}

func New(authData AuthData) (YaClAPI, error) {

	var creds ycsdk.Credentials
	var err error
	var yapi = YaClAPI{}

	switch authData.Type {
	case "oauth":
		token, exists := os.LookupEnv("OAUTH_TOKEN")
		if !exists {
			return yapi, errors.New("отсутствует OAUTH_TOKEN в окружении")
		}

		creds, err = NewOAuthCreds(token)

	case "instance":
		creds, err = NewInstanceCreds()

	case "service_account":
		creds, err = NewServiceAccountCreds(authData.ServiceAccountKeyPath)

	default:
		return yapi, errors.New("неизвестный способ авторизации: " + authData.Type)
	}

	yapi.Ctx = context.Background()
	sdkInstance, err := ycsdk.Build(yapi.Ctx, ycsdk.Config{
		Credentials: creds,
	})
	if err != nil {
		return yapi, err
	}

	yapi.Sdk = sdkInstance
	log.Println("Успешная инициализация Yandex Cloud API")

	return yapi, nil
}
