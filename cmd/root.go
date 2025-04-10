package cmd

import (
	"context"
	"github.com/urfave/cli/v3"
	"log"
	"os"
	"ysd-cli/api"
)

func Execute() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			deployCmd,
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var Environment string
var FuncDir string
var Functions []function.Function
var YC api.YaClAPI

func init() {
	//cobra.OnInitialize(initCli)
	//if err := godotenv.Load(); err != nil {
	//	log.Print("No .env file found")
	//}

	//rootCmd.PersistentFlags().StringVar(&Environment, "env", "production", "Окружение (production, dev и т.д.)")
	//rootCmd.AddCommand(deployCmd)

}

func initCli() {
	FuncDir = config.MainConfig.BaseDir
	Functions = function.GetFunctionList(FuncDir, Environment)

	log.Println("Использование окружения:", Environment)

	initApi()

	for i, f := range Functions {
		apiData, err := YC.GetFunction(f.Config.Id)
		if err != nil {
			log.Println("Ошибка получения данных функции", f.PathName+":")
			log.Println(err)
		}
		Functions[i].ApiData = apiData
	}
}
