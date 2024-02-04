package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/bruttazz/devc/internal"
)

// copy a file from source to destination
func copyFile(dst, src string) (int64, error) {
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// check if cache exist or not
func CheckCacheExists(file string) (exist bool) {
	fileInfo, err := os.Stat(file)
	// exist = !errors.Is(err, os.ErrNotExist)
	if err == nil {
		exist = fileInfo.ModTime().Sub(time.Now()).Hours() < float64(internal.Config.CacheExpiryHrs)
	}
	return
}

func GetCache(dest string, file string) (err error) {
	srcFile := filepath.Join(internal.Config.CacheDir, file)

	exist := CheckCacheExists(srcFile)
	if !exist {
		return errors.New(fmt.Sprintf("cache not found (%v)", srcFile))
	}
	_, err = copyFile(dest, srcFile)
	return
}

func SetCache(src string, name string, subdir string) (err error) {
	os.MkdirAll(filepath.Join(internal.Config.CacheDir, subdir), 0755)
	_, err = copyFile(filepath.Join(internal.Config.CacheDir, subdir, name), src)
	return
}
