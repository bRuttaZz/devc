package cmd

import (
	"github.com/bruttazz/devc/internal/builder"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var logoutOptions configs.LogoutCmdOptions

var logoutCmd = &cobra.Command{
	Use:   "logout registry-uri",
	Short: "logout from a previously logged in registry",
	Long:  `logout from a previously logged in registries using "devc login".`,
	Run: func(cmd *cobra.Command, args []string) {
		builder.Logout(&logoutOptions, args)
	},
}
