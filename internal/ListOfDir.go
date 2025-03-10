package internal

import (
	"errors"
	"log"
	"os"
	"ycf-cli/config"
)

func ListOfDir(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var dirs []string
	for _, e := range entries {
		fileName := e.Name()
		if e.IsDir() && fileName[0:1] != "." {
			if _, err := os.Stat(dir + "/" + fileName + "/" + config.CONFIG_FILE); !errors.Is(err, os.ErrNotExist) {
				dirs = append(dirs, e.Name())
			}
		}
	}
	return dirs
}
