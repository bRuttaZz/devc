package builder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bruttazz/devc/internal/configs"
)

func runCommand(name string, cmd []string) (err error) {
	cmd_ := exec.Command(
		name,
		cmd...,
	)
	cmd_.Stderr = os.Stderr
	cmd_.Stdout = os.Stdout
	cmd_.Stdin = os.Stdin
	err = cmd_.Run()
	return
}

// provide global options for buildah
// args:
//
//	envPath - env path
func getGlobalBuildahOptions(abs, envPath string) (cmd []string) {
	var rootPath string
	if len(envPath) > 0 {
		rootPath = filepath.Join(abs, envPath, configs.Config.EnvSettings.BuildDir)
	} else {
		rootPath = filepath.Join(abs, configs.Config.CacheDir, configs.Config.CacheDirSettings.Buildah)
	}
	cmd = []string{
		"--root",
		rootPath,
		"--storage-driver",
		configs.Config.CacheDirSettings.StorageDriver,
	}
	return
}

// get bulid command options
// args:
//
//	envPath - env path
func getBuildOptions(abs, envPath string) (cmd []string) {
	cmd = []string{
		"build",
		"--rm",
		"--layers=false",
		fmt.Sprintf(
			"--output=type=local,dest=%v",
			filepath.Join(abs, envPath, configs.Config.EnvSettings.RootDir),
		),
	}
	return
}
