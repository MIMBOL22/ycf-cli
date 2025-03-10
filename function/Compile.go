package function

import (
	"log"
	"os"
	"os/exec"
)

func (f *Function) Compile() {
	_, entryDir, funcCacheDir, _ := f.getPaths()

	log.Println("TSC сборка функции", f.PathName)

	cmd := exec.Command("tsc", "--outDir", "../../"+funcCacheDir+"/tsc", "-p", "./tsconfig.json")
	cmd.Dir = entryDir
	out, err := cmd.Output()

	if err != nil {
		log.Println("Ошибка TSC сборки функции", f.PathName, ":\n", string(out))
		os.Exit(1)
	}
}
