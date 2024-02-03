package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	// "github.com/containers/buildah"
)

var rootCmd = &cobra.Command{
	Use:   "devc [command] [options] [name/url]",
	Short: "Containers for Developers (Container as a Directory)",
	Long: `devc: Containers for Developers (Container as a Directory)!
	Provides with a developer friendly interface to build your application with containers!`,
}

func Execute() {
	// fmt.Println("testing %T", )
	err := rootCmd.Execute()
	if err != nil {
		rootCmd.SilenceUsage = false
		rootCmd.SilenceErrors = false

		fmt.Println("[Error] ", err)
		os.Exit(1)
	}
}
