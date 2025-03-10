package function

import "ycf-cli/config"

func (f *Function) getPaths() (string, string, string, string) {
	entryDir := config.Config.BaseDir + "/" + f.PathName
	entryFile := entryDir + "/" + f.Config.Entrypoint.File
	funcCacheDir := "./" + config.CACHE_DIR + "/cache/" + f.PathName
	outDir := funcCacheDir + "/dist"
	return entryFile, entryDir, funcCacheDir, outDir
}
