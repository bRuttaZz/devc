package main

import (
	"fmt"
	"os"

	"github.com/bruttazz/devc/cmd"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/bruttazz/devc/internal/env"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("\n[Error] : %v\n", v)
			os.Exit(1)
		}
	}()
	configs.LoadConfig()

	err := env.SetupEnv("build")
	if err != nil {
		panic(err.Error())
	}
	cmd.Execute()
}
