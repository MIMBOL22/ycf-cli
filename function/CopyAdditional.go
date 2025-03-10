package function

import (
	"errors"
	"log"
	"os"
	"strings"
	"sync"
	"ycf-cli/internal"
)

func (f *Function) CopyAdditional() {
	_, entryDir, funcCacheDir, _ := f.getPaths()

	log.Println("Копирование дополнительных файлов фунации", f.PathName)

	var wg sync.WaitGroup
	for _, file := range f.Config.AdditionalFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			filePath := funcCacheDir + "/dist/" + file

			// Создание структуры папок для копируемого файла
			parts := strings.Split(filePath, "/")
			onlyDirParts := parts[0 : len(parts)-1]
			onlyDir := strings.Join(onlyDirParts, "/")
			if _, err := os.Stat(onlyDir); errors.Is(err, os.ErrNotExist) {
				err := os.MkdirAll(onlyDir, 0775)
				if err != nil {
					log.Println("Ошибка создания древа для доп. файлов функции", f.PathName, ":")
					log.Println(err)
					os.Exit(1)
				}
			}

			_, err := internal.CopyFile(entryDir+"/"+file, filePath)
			if err != nil {
				log.Println("Ошибка копирования доп. файлов функции", f.PathName, ":")
				log.Println(err)
				os.Exit(1)
			}
		}(file)
	}
	wg.Wait()
}
