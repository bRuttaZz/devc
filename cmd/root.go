package cmd

import (
	"fmt"
	"os"

	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
	// "github.com/containers/buildah"
)

var roomCmdOptions configs.RootOptions

var rootCmd = &cobra.Command{
	Use:   "devc",
	Short: "Containers for Developers (Container as a Directory)",
	Long: `devc: Containers for Developers (Container as a Directory)!
	Provides with a developer friendly interface to build your application with containers!`,
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&roomCmdOptions.Verbose, "verbose", false, "Dispaly detailed debugging logs")

	rootCmd.AddCommand(activateCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(imagesCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(rmiCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(pruneCmd)
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		rootCmd.SilenceUsage = false
		rootCmd.SilenceErrors = false

		fmt.Println("[Error] ", err)
		os.Exit(1)
	}
}
