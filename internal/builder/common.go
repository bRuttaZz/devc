package builder

import (
	// "fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bruttazz/devc/internal/configs"
	"github.com/bruttazz/devc/internal/utils"
)

// execute command
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
func getGlobalBuildahOptions() (cmd []string) {
	var rootPath string
	rootPath = filepath.Join(configs.Config.CacheDir, configs.Config.CacheDirSettings.BuildahCache)
	cmd = []string{
		"--root",
		rootPath,
		"--storage-driver",
		configs.Config.CacheDirSettings.StorageDriver,
	}
	return
}

// get authfile option for buildah cli
func getAuthFileOptions() (cmd []string, err error) {
	var authFile = filepath.Join(
		configs.Config.CacheDir,
		configs.Config.CacheDirSettings.LoginAuthFile,
	)
	cmd = []string{"--authfile", authFile}
	err = utils.TouchAJSONFile(authFile)
	return
}

// get bulid command options
func getBuildOptions() (imageName string, cmd []string, err error) {
	var authFileOptions []string
	authFileOptions, err = getAuthFileOptions()
	imageName = utils.CreateRandomString()
	cmd = []string{
		"build",
		"--force-rm",
		"--rm",
		"--layers=false",
		"--tag",
		imageName,
		authFileOptions[0],
		authFileOptions[1],
	}
	return
}

// garbage collection
func clearBuildCache() (err error) {
	var options = getGlobalBuildahOptions()
	options = append(options, "prune")
	options = append(options, "-af")
	err = runCommand(configs.Config.Buildah.Path, options)
	if err != nil {
		return
	}
	err = runCommand("rm", []string{"-rf", options[1]})
	return
}
