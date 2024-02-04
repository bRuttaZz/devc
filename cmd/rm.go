package cmd

import (
	"github.com/bruttazz/devc/internal/builder"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var rmOptions configs.RmCmdOptions

var rmCmd = &cobra.Command{
	Use:   "rm env-name",
	Short: "remove a devc environment",
	Long: `remove the specified devc environment. 
Can be used if it's not able to remove the directory otherwise. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		builder.Rm(&rmOptions, args)
	},
}
