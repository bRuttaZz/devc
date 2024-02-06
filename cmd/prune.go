package cmd

import (
	"github.com/bruttazz/devc/internal/builder"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var pruneOptions configs.PruneCmdOptions

var pruneCmd = &cobra.Command{
	Use:   "prune [flags]",
	Short: "prune all the cached images.",
	Long: `Remove all the cached images. 
Images will be cached for latter use if devc env is created using "pull" command.
Using --wipe option will clear everything including the registry login credentials
`,
	Run: func(cmd *cobra.Command, args []string) {
		builder.Prune(&pruneOptions, args)
	},
}

func init() {
	pruneCmd.PersistentFlags().BoolVar(&pruneOptions.Wipe, "wipe", false, "Clear all cache from user session, including registry login credentials!")
}
