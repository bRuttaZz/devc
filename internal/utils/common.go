package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"math/rand"
)

// for random number generation
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// generate random string of fixed length
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// generate a uuid / or standard sized string
func CreateRandomString() (outString string) {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		fmt.Println("[warning] error generating uuid : ", err)
		fmt.Println("[warning] forwarding with random string generation!")
	} else {
		outString = strings.Trim(string(newUUID), "\n")
	}
	if len(outString) != 36 {
		outString = RandString(36)
	}
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
