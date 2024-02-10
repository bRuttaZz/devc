package internal

import (
	"os/exec"

	"github.com/bruttazz/devc/internal/configs"
)

func TestBuildah() bool {
	cmd := exec.Command(configs.Config.Buildah.Path, "-v")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
