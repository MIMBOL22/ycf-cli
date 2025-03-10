package api

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
	"log"
	"os"
)

func (y *YaClAPI) GetVersions(funId string, filter string) *functions.ListFunctionsVersionsResponse {
	versions, err := y.Sdk.Serverless().Functions().Function().ListVersions(y.Ctx, &functions.ListFunctionsVersionsRequest{
		Id: &functions.ListFunctionsVersionsRequest_FunctionId{
			FunctionId: funId,
		},
		Filter: filter,
	})
	if err != nil {
		log.Println("Ошибка получения версий функции с ID", funId, ":")
		log.Println(err)
		os.Exit(1)
	}

	return versions
}
