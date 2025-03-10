package api

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
	"log"
	"os"
)

func (y *YaClAPI) RemoveVersion(versionId string) {
	_, err := y.Sdk.Serverless().Functions().Function().DeleteVersion(y.Ctx, &functions.DeleteFunctionVersionRequest{
		FunctionVersionId: versionId,
	})
	if err != nil {
		log.Println("Ошибка удаления версии функции", versionId, ":")
		log.Println(err)
		os.Exit(1)
	}
}
