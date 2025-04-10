package utils

import (
	"errors"
	"os"
)

func ListOfDir(dir, functionConfigFilename string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return make([]string, 0), err
	}

	var dirs []string
	for _, e := range entries {
		fileName := e.Name()
		if e.IsDir() && fileName[0:1] != "." {
			if _, err := os.Stat(dir + "/" + fileName + "/" + functionConfigFilename); !errors.Is(err, os.ErrNotExist) {
				dirs = append(dirs, e.Name())
			}
		}
	}
	return dirs, nil
}
