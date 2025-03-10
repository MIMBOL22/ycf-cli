package function

import (
	"log"
	"os"
	"ycf-cli/internal"
)

func (f *Function) Zip() {
	_, _, funcCacheDir, _ := f.getPaths()

	err := internal.Zip(funcCacheDir+"/dist", funcCacheDir+"/func.zip")
	if err != nil {
		log.Println("Ошибка архивации функции", f.PathName, ":")
		log.Println(err)
		os.Exit(1)
	}
}
