package utils

import (
	"os"
)

// make a file executable
func MakeExecutable(path string) (err error) {
	info, err := os.Stat(path)
	if err == nil {
		newMod := info.Mode() | 0100
		err = os.Chmod(path, newMod)
	}
	return
}
