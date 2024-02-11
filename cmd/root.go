package cmd

import (
	"fmt"
	"os"

	"github.com/bruttazz/devc/internal"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var roomCmdOptions configs.RootOptions

var rootCmd = &cobra.Command{
	Use:   "devc",
	Short: "Containers for Developers (Container as a Directory)",
	Long: `devc: Containers for Developers (Container as a Directory)!
Offers developers a user-friendly interface for constructing applications
using containers! Generate effortlessly usable virtual environments from
Dockerfiles, other Containerfiles or container images, activate them, 
engage in development, implement changes, and reuse seamlessly.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if roomCmdOptions.Version {
			err := internal.TestBuildah()
			if !err {
				panic(fmt.Sprintf("cannot resolve dependency : %v", configs.Config.Buildah.Path))
			}
			fmt.Printf("devc %v\n", configs.Config.Version)
			return
		}
		fmt.Printf("\nDEVC %v\n\n", configs.Config.Version)
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&roomCmdOptions.Version, "version", "v", false, "Get current version")

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
