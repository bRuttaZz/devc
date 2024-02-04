package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout from a registry",
	Long:  `logout from a registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		// print a nice version with options
		fmt.Println("TO BE IMPLEMENTED!")
	},
}
