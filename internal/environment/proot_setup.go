package environment

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bruttazz/devc/internal"
)

// For downloading Proot setup
// TODO : add download progress
func DownloadProot(out_file_path string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	file, err := os.Create(out_file_path)
	if err != nil {
		panic("error creating proot binary: " + err.Error())
	}
	defer file.Close()

	resp, err := http.Get(internal.Config.Proot.Url)
	if err != nil {
		panic("error downloading proot binary: " + err.Error())
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic("error forming proot binary: " + err.Error())
	}

	return nil
}
