package api

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
	"google.golang.org/protobuf/types/known/durationpb"
	"log"
	"os"
	ycfunc "ycf-cli/function"
	"ycf-cli/internal"
)

func (y *YaClAPI) CreateVersion(fun *ycfunc.Function) {
	log.Println("Загрузка функции", fun.PathName, "в облако")

	fConfig := fun.Config
	entrypoint := internal.GetFilename(fConfig.Entrypoint.File) + "." + fConfig.Entrypoint.Function

	_, err := y.Sdk.Serverless().Functions().Function().CreateVersion(y.Ctx, &functions.CreateFunctionVersionRequest{
		FunctionId:  fConfig.Id,
		Runtime:     fConfig.Runtime,
		Description: fConfig.Description,
		Entrypoint:  entrypoint,
		Resources: &functions.Resources{
			Memory: fConfig.Memory * 1024 * 1024,
		},
		ExecutionTimeout: &durationpb.Duration{
			Seconds: fConfig.Timeout,
		},
		ServiceAccountId: fConfig.ServiceAccountId,
		PackageSource: &functions.CreateFunctionVersionRequest_Package{
			Package: fun.S3Package,
		},
		Environment:   nil,
		Connectivity:  nil,
		Secrets:       nil,
		LogOptions:    nil,
		StorageMounts: nil,
		Mounts:        nil,
	})
	if err != nil {
		log.Println("Ошибка загрузки функции", fun.PathName, "в облако:")
		log.Println(err)
		os.Exit(1)
	}
}
