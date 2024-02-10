package main

import (
	"fmt"
	"os"

	"github.com/bruttazz/devc/cmd"
	"github.com/bruttazz/devc/internal/configs"
)

var version string = ""

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("\n[Error] : %v\n", v)
			os.Exit(1)
		}
	}()
	configs.LoadConfig()
	configs.Config.Version = version
	cmd.Execute()
}
