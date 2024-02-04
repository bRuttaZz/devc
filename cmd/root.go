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
	Run: func(cmd *cobra.Command, args []string) {
		// print a nice version with options
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&roomCmdOptions.Verbose, "verbose", false, "Dispaly detailed debugging logs")

	rootCmd.AddCommand(activateCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(psCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(rmiCmd)
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
