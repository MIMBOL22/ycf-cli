package cmd

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
	"ycf-cli/api"
	"ycf-cli/config"
	"ycf-cli/function"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ycf-cli-src",
	Short: "Cli for build and deploy functions to Yandex Cloud",
	Long:  `Cli for build and deploy functions to Yandex Cloud`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome to ycf-cli-src!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var Environment string
var FuncDir string
var Functions []function.Function
var YC api.YaClAPI

func initApi() {
	YC = api.YaClAPI{}

	switch config.Config.AuthType {
	case "oauth":
		token, exists := os.LookupEnv("OAUTH_TOKEN")
		if !exists {
			log.Println("Отсутствует OAUTH_TOKEN в окружении")
			os.Exit(1)
		}
		YC.OAuthAuth(token)

	case "instance":
		YC.InstanceAuth()

	case "service_account":
		YC.ServiceAuth(config.Config.ServiceAccountKeyPath)

	default:
		log.Println("Неизвестный способ авторизации:", config.Config.AuthType)
		log.Println("Поддерживаемые способы:", "oauth", "instance", "service_account")
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initCli)
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	rootCmd.PersistentFlags().StringVar(&Environment, "env", "production", "Окружение (production, dev и т.д.)")
	rootCmd.AddCommand(deployCmd)

}

func initCli() {
	FuncDir = config.Config.BaseDir
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
