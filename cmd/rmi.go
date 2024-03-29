package cmd

import (
	"github.com/bruttazz/devc/internal/builder"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var rmiOptions configs.RmiCmdOptions

var rmiCmd = &cobra.Command{
	Use:   "rmi image-id",
	Short: "remove a cached image",
	Long: `remove a cached image by id
Images will be cached for latter use if devc env is created using "pull" command`,
	Run: func(cmd *cobra.Command, args []string) {
		builder.Rmi(&rmiOptions, args)
	},
}
