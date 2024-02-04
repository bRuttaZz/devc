package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "activate a devc.",
	Long: `activate a devc environement. On activating it will automatically mount current working
	dir to an unused directory inside the environment`,
	Run: func(cmd *cobra.Command, args []string) {
		// print a nice version with options
		fmt.Println("got here on activate")
	},
}
