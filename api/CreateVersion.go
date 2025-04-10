package api

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
	"google.golang.org/protobuf/types/known/durationpb"
	"log"
	"ysd-cli/utils"
)

type CreateVersionData struct {
	id               string
	runtime          string
	description      string
	entrypoint       string
	memory           int64
	timeout          int64
	serviceAccountId string
	s3Package        *functions.Package
	fileName         string
	functionCodeName string
	pathName         string
}

func (y *YaClAPI) CreateVersion(data *CreateVersionData) error {
	log.Println("Загрузка функции", data.pathName, "в облако")

	entrypoint := utils.GetFilename(data.fileName) + "." + data.functionCodeName

	_, err := y.Sdk.Serverless().Functions().Function().CreateVersion(y.Ctx, &functions.CreateFunctionVersionRequest{
		FunctionId:  data.id,
		Runtime:     data.runtime,
		Description: data.description,
		Entrypoint:  entrypoint,
		Resources: &functions.Resources{
			Memory: data.memory * 1024 * 1024,
		},
		ExecutionTimeout: &durationpb.Duration{
			Seconds: data.timeout,
		},
		ServiceAccountId: data.serviceAccountId,
		PackageSource: &functions.CreateFunctionVersionRequest_Package{
			Package: data.s3Package,
		},
		Environment:   nil,
		Connectivity:  nil,
		Secrets:       nil,
		LogOptions:    nil,
		StorageMounts: nil,
		Mounts:        nil,
	})
	return err
}
