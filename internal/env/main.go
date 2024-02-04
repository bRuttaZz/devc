package env

import (
	"os"
	"path/filepath"

	"github.com/bruttazz/devc/internal/configs"
)

// setup devc bin and add namespace and other modifications to env dir
func SetupEnv(envPath string) (err error) {
	// create bin dir
	os.MkdirAll(filepath.Join(envPath, configs.Config.EnvSettings.DevcBin), 0755)
	err = setProot(envPath)
	if err != nil {
		return
	}

	// setup activate script
	err = setupActivateScript(envPath)
	if err != nil {
		return
	}

	// modifications to the container dir
	err = finishUpRootBin(envPath)

	return
}
