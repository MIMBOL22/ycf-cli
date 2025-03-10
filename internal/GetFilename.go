package internal

import (
	"regexp"
)

func GetFilename(fullPath string) string {
	re, _ := regexp.Compile(`/([a-zA-Z0-9_-]+)\.(t|j)s`)
	res := re.FindStringSubmatch(fullPath)

	return res[1]
}
