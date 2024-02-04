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

func WriteTextToFile(path string, text string) (err error) {
	fileObj, err := os.Create(path)
	if err != nil {
		return
	}
	defer fileObj.Close()
	_, err = fileObj.Write([]byte(text))
	return
}
