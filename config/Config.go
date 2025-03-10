package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Config MonorepConfig

func init() {
	Config = GetMonorepConfig()
}

type MonorepConfig struct {
	BaseDir               string `yaml:"base_dir"`
	AuthType              string `yaml:"auth_type"`
	ServiceAccountKeyPath string `yaml:"service_account_key_path"`
	S3Bucket              string `yaml:"s3_bucket"`
}

func GetMonorepConfig() MonorepConfig {
	yamlFile, err := os.ReadFile(MONOREP_FILE)
	if err != nil {
		log.Printf("Err to get monorep config:  #%v ", err)
	}

	var c MonorepConfig
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
