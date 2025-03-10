package internal

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Zip(inDir, outZip string) error {
	file, err := os.Create(outZip)
	if err != nil {
		log.Println("Ошибка создания файла архива", outZip, ":")
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Замена путей для сохранения относительности
		replaceFrom := strings.Replace(inDir, "./", "", 1) + "/"
		newPath := strings.Replace(path, replaceFrom, "", -1)

		// Сохранение архива
		f, err := w.Create(newPath)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}
		return nil
	}

	return filepath.Walk(inDir, walker)
}
