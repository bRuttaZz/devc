package environment

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bruttazz/devc/internal"
	"github.com/bruttazz/devc/internal/utils"
)

func SetProot(envPath string) (err error) {
	outfile := filepath.Join(envPath, "bin", "proot")

	err = utils.GetCache(outfile, "proot")
	if err != nil {

		fmt.Println("[proot setup] no cache found! attempt downloading..", err)
		err = utils.DownloadFile(outfile, internal.Config.Proot.Url)
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

func SetupBin(envPath string) (err error) {
	// create bin dir
	os.MkdirAll(filepath.Join(envPath, "bin"), 0755)
	// setup proot
	err = SetProot(envPath)
	return
}
