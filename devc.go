/*
devc: Containers for Developers (Container as a Directory)!
(command line utility)

Offers developers a user-friendly interface for constructing applications
using containers! Generate reusable usable virtual environments from
Containerfiles or container images, activate them, engage in development,
implement changes, and reuse seamlessly
*/
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
