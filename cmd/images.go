package cmd

import (
	"github.com/bruttazz/devc/internal/builder"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var imagesOptions configs.ImagesCmdOptions

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "show all the cached images",
	Long: `Show all the cache images. 
Images will be cached for latter use if devc env is created using "pull" command`,
	Run: func(cmd *cobra.Command, args []string) {
		builder.Images(&imagesOptions, args)
	},
}
