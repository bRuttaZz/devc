package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a devc environment",
	Long:  `remove a devc encironment`,
	Run: func(cmd *cobra.Command, args []string) {
		// print a nice version with options
		fmt.Println("got here on rm")
	},
}
