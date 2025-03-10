package function

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"ycf-cli/config"
	"ycf-cli/internal"
)

func GetFunctionList(funcBaseDir, environment string) []Function {
	dirs := internal.ListOfDir(funcBaseDir)
	funcConfigArray := make([]Function, len(dirs))

	for i, funcDir := range dirs {
		yamlFile, err := os.ReadFile(funcBaseDir + "/" + funcDir + "/" + config.CONFIG_FILE)
		if err != nil {
			log.Printf("Err to get config:  #%v ", err)
		}

		var functionEnvironments FunctionEnvironments
		err = yaml.Unmarshal(yamlFile, &functionEnvironments)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}

		for _, fEnvironment := range functionEnvironments.Environments {
			if fEnvironment.EnvironmentName == environment {
				funcConfigArray[i].Config = fEnvironment
				break
			}
		}
		funcConfigArray[i].PathName = funcDir
	}

	return funcConfigArray
}
