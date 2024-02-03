package main

import (
	"fmt"
	"os"

	"github.com/bruttazz/devc/cmd"
	"github.com/bruttazz/devc/internal"
	"github.com/bruttazz/devc/internal/environment"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("\n[Error] : %v\n", v)
			os.Exit(1)
		}
	}()
	internal.LoadConfig()

	err := environment.DownloadProot("build/proot")
	if err != nil {
		panic(err.Error())
	}
	cmd.Execute()
}
