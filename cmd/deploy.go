package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
	"log"
	"os"
	"sync"
	"ycf-cli/function"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy functions to YCloud",
	Run: func(cmd *cobra.Command, args []string) {

		if len(Functions) == 0 {
			log.Println("Функции для деплоя не найдены")
			os.Exit(0)
		}

		// Компиляция, сборка, архивирование
		var wg sync.WaitGroup
		for _, f := range Functions {
			wg.Add(1)
			go func(f function.Function) {
				defer wg.Done()

				f.Build()
				f.Zip()
			}(f)
		}
		wg.Wait()

		// Загрузка в S3
		for i, _ := range Functions {
			wg.Add(1)
			go func(f *function.Function) {
				defer wg.Done()
				f.Upload2S3()
			}(&Functions[i])
		}
		wg.Wait()

		// Загрузка в облако
		for i, _ := range Functions {
			wg.Add(1)
			go func(f *function.Function) {
				defer wg.Done()

				YC.CreateVersion(f)
			}(&Functions[i])
		}
		wg.Wait()

		// Удаление прошлых версий
		for i, _ := range Functions {
			wg.Add(1)
			go func(f *function.Function) {
				defer wg.Done()

				versions := YC.GetVersions(f.Config.Id, `status="ACTIVE"`)
				for _, v := range versions.Versions {
					wg.Add(1)
					go func(v *functions.Version) {
						defer wg.Done()

						if len(v.Tags) == 0 {
							YC.RemoveVersion(v.Id)
						}
					}(v)
				}
			}(&Functions[i])
		}
		wg.Wait()

		log.Println("Успешный деплой всех функций!")
	},
}
