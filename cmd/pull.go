package cmd

import (
	"github.com/bruttazz/devc/internal/builder"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var pullOptions configs.PullCmdOptions

var pullCmd = &cobra.Command{
	Use:   "pull [options] image-name env-name",
	Short: "build a devc env directly from a container image",
	Long: `build an devc environment from a container image.
The first argument 'image-name' (optionally with a tag) will be search inside
available registries. One can add new registries by executing "devc login" 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		builder.Pull(&pullOptions, args)
	},
}
