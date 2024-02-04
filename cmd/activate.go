package cmd

import (
	"github.com/bruttazz/devc/internal/builder"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var activateOptions configs.ActivateCmdOptions

var activateCmd = &cobra.Command{
	Use:   "activate env-name",
	Short: "activate a devc env",
	Long: `activate a devc environement. 
One can also enable the env by executing "source <env-name>/bin/activate".
On activating the env, it will automatically mount the current working
dir to an unused directory (/devc) inside the environment`,
	Run: func(cmd *cobra.Command, args []string) {
		builder.Activate(&activateOptions, args)
	},
}

func init() {
	activateCmd.PersistentFlags().StringVarP(&activateOptions.Workdir, "working-dir", "w", "", "Use a deferent working directory on activation, defaults to '/devc'")
}
