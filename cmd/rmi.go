package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "remove a cached image",
	Long:  `remove a cached image`,
	Run: func(cmd *cobra.Command, args []string) {
		// print a nice version with options
		fmt.Println("got here on rmi")
	},
}
