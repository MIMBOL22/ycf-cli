package api

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
)

func (y *YaClAPI) RemoveVersion(versionId string) error {
	_, err := y.Sdk.Serverless().Functions().Function().DeleteVersion(y.Ctx, &functions.DeleteFunctionVersionRequest{
		FunctionVersionId: versionId,
	})
	return err
}
