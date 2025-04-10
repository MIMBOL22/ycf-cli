package api

import (
	"encoding/json"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"
	"os"
)

// TODO: Протестировать
func NewInstanceCreds() (ycsdk.Credentials, error) {
	return ycsdk.InstanceServiceAccount(), nil
}

func NewServiceAccountCreds(keyFilePath string) (ycsdk.Credentials, error) {
	keyFile, err := os.ReadFile(keyFilePath)
	if err != nil {
		//log.Fatalf("Ошибка при чтении файла ключа: %v", err)
		return nil, err
	}

	var key *iamkey.Key
	if err := json.Unmarshal(keyFile, &key); err != nil {
		//log.Fatalf("Ошибка при обработке JSON ключа: %v", err)
		return nil, err
	}

	credentials, err := ycsdk.ServiceAccountKey(key)
	if err != nil {
		//log.Fatalf("Ошибка авторизации через ServiceAccount: %v", err)
		return nil, err
	}

	return credentials, nil
}

func NewOAuthCreds(oAuthToken string) (ycsdk.Credentials, error) {
	return ycsdk.OAuthToken(oAuthToken), nil
}
