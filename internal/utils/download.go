package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// print download percent in std out
func printDownloadPercent(done chan bool, path string, total int64) {
	var stop bool
	var status bool

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("\n[WARNING] error showing download progress : %v\n", err)
		return
	}
	defer file.Close()

	for {
		select {
		case status = <-done:
			stop = true
		default:
			fi, err := file.Stat()
			if err != nil {
				fmt.Printf("\n[WARNING] error showing download progress : %v\n", err)
				return
			}

			size := fi.Size()
			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100

			fmt.Printf("\rdownloading : [%v] %.0f%%", path, percent)
		}

		if stop {
			if status {
				fmt.Printf("\rdownloading : [%v] 100%%\n", path)
			} else {
				fmt.Printf("\r[Error] downloading error : (%v)\n", path)
			}
			break
		}

		time.Sleep(time.Second)
	}
}

// Download a file from given url to specified out setup
// TODO : add download progress setup
func DownloadFile(out_path string, url string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	file, err := os.Create(out_path)
	if err != nil {
		panic("error creating proot binary: " + err.Error())
	}
	defer file.Close()

	// get file size and setup progress printing
	headResp, err := http.Head(url)
	if err != nil {
		panic("error downloading proot binary: " + err.Error())
	}
	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	if err != nil {
		panic("error preparing proot download: " + err.Error())
	}
	done := make(chan bool)
	go printDownloadPercent(done, out_path, int64(size))

	// start downloading
	resp, err := http.Get(url)
	if err != nil {
		done <- false
		panic("error downloading proot binary: " + err.Error())
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		done <- false
		panic("error forming proot binary: " + err.Error())
	}
	done <- true

	return
}
