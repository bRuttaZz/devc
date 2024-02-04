package main

import (
	"fmt"
	"os"

	"github.com/bruttazz/devc/cmd"
	"github.com/bruttazz/devc/internal/configs"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("\n[Error] : %v\n", v)
			os.Exit(1)
		}
	}()
	configs.LoadConfig()
	cmd.Execute()
}
