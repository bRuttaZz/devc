package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login into a registry",
	Long:  `login into a registry. using the username and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		// print a nice version with options
		fmt.Println("got here on login")
	},
}
