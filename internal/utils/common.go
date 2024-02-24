package utils

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// generate a uuid / or standard sized string
func CreateRandomString() (outString string) {
	outString = uuid.New().String()
	return
}

// make a file executable
func MakeExecutable(path string) (err error) {
	info, err := os.Stat(path)
	if err == nil {
		newMod := info.Mode() | 0111
		err = os.Chmod(path, newMod)
	}
	return
}

// write text to a file
func WriteTextToFile(path string, text string) (err error) {
	fileObj, err := os.Create(path)
	if err != nil {
		return
	}
	defer fileObj.Close()
	_, err = fileObj.Write([]byte(text))
	return
}

// touch a file (create if not exists, otherwise do nothing (not a complete UNIX touch))
func TouchAJSONFile(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(path), 0755)
		return WriteTextToFile(path, "{}")

	}
	return err
}
