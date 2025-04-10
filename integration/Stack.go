package integration

import (
	"gopkg.in/yaml.v3"
	"os"
	"ysd-cli/api"
)

type Stack struct {
	Version      int8          `yaml:"version"`
	BucketName   string        `yaml:"bucket_name"`
	Auth         api.AuthData  `yaml:"auth"`
	Integrations []Integration `yaml:"integrations"`
}

func NewStack() (Stack, error) {
	var s Stack
	yamlFile, err := os.ReadFile(MonorepFile)

	if err != nil {
		return s, err
	}
	err = yaml.Unmarshal(yamlFile, &s)

	return s, err
}
