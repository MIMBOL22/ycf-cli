package function

import (
	ysdkFunc "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
)

type FunctionConfigEnvironment struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Value  string `yaml:"value"`
	Secret struct {
		Id        string `yaml:"id"`
		VersionId string `yaml:"version_id"`
		Key       string `yaml:"key"`
	} `yaml:"secret"`
}

type FunctionConfigEntryPoint struct {
	File     string `yaml:"file"`
	Function string `yaml:"function"`
}

type FunctionConfig struct {
	EnvironmentName    string                      `yaml:"name"`
	Id                 string                      `yaml:"id"`
	Memory             int64                       `yaml:"memory"`
	Runtime            string                      `yaml:"runtime"`
	Description        string                      `yaml:"description"`
	Entrypoint         FunctionConfigEntryPoint    `yaml:"entrypoint"`
	EntrypointFunction string                      `yaml:"entrypoint_function"`
	Timeout            int64                       `yaml:"timeout"`
	ServiceAccountId   string                      `yaml:"service_account_id"`
	Environment        []FunctionConfigEnvironment `yaml:"environment"`
	AdditionalFiles    []string                    `yaml:"additional_files"`
}

type Function struct {
	Config    FunctionConfig
	PathName  string
	ApiData   *ysdkFunc.Function
	S3Package *ysdkFunc.Package
}

type FunctionEnvironments struct {
	Environments []FunctionConfig `yaml:"environments"`
}
