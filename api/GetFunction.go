package api

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
)

func (y *YaClAPI) GetFunction(id string) (*functions.Function, error) {
	return y.Sdk.Serverless().Functions().Function().Get(y.Ctx, &functions.GetFunctionRequest{
		FunctionId: id,
	})
}
