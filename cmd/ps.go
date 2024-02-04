package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "show all the cached images",
	Long:  `show all the cache images`,
	Run: func(cmd *cobra.Command, args []string) {
		// print a nice version with options
		fmt.Println("got here on ps")
	},
}
