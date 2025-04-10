package api

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
)

func (y *YaClAPI) GetVersions(funId string, filter string) (*functions.ListFunctionsVersionsResponse, error) {
	return y.Sdk.Serverless().Functions().Function().ListVersions(y.Ctx, &functions.ListFunctionsVersionsRequest{
		Id: &functions.ListFunctionsVersionsRequest_FunctionId{
			FunctionId: funId,
		},
		Filter: filter,
	})
}
