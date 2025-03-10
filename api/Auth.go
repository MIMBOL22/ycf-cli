package api

import (
	"encoding/json"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"
	"log"
	"os"
)

// TODO: Протестировать
func (y *YaClAPI) InstanceAuth() {
	y.Init(ycsdk.InstanceServiceAccount())
}

func (y *YaClAPI) ServiceAuth(keyFilePath string) {
	keyFile, err := os.ReadFile(keyFilePath)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла ключа: %v", err)
	}

	var key *iamkey.Key
	if err := json.Unmarshal(keyFile, &key); err != nil {
		log.Fatalf("Ошибка при обработке JSON ключа: %v", err)
	}

	credentials, err := ycsdk.ServiceAccountKey(key)
	if err != nil {
		log.Fatalf("Ошибка авторизации через ServiceAccount: %v", err)
	}

	y.Init(credentials)
}

func (y *YaClAPI) OAuthAuth(oAuthToken string) {
	y.Init(ycsdk.OAuthToken(oAuthToken))
}
