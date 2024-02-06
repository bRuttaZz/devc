package env

import (
	"fmt"
	"path/filepath"

	"github.com/bruttazz/devc/internal/configs"
	"github.com/bruttazz/devc/internal/utils"
)

func setProot(envPath string) (err error) {
	outfile := filepath.Join(envPath, configs.Config.EnvSettings.DevcBin, "proot")

	err = utils.GetCache(outfile, configs.Config.CacheDirSettings.ProotCache)
	if err != nil {

		fmt.Println("[proot setup] no cache found! attempt downloading..", err)
		err = utils.DownloadFile(outfile, configs.Config.Proot.Url)
		if err == nil {
			e := utils.SetCache(outfile, "proot", "")
			if e != nil {
				fmt.Printf("[proot setup] error setting cache : %v ! skipping..", e)
			}
		}
	}
	if err == nil {
		err = utils.MakeExecutable(outfile)
	}
	return
}
