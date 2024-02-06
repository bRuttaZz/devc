package cmd

import (
	"github.com/bruttazz/devc/internal/builder"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var loginOptions configs.LoginCmdOptions

var loginCmd = &cobra.Command{
	Use:   "login registry-uri",
	Short: "login into a container-image registry",
	Long: `login into a container-image registry, using the username and password.
Can be use to pull images from private registries. "registry-uri" can be the domain name 
or domain with port number. eg: 'registry.hub.docker.com'    
	`,
	Run: func(cmd *cobra.Command, args []string) {
		builder.Login(&loginOptions, args)
	},
}

func init() {
	loginCmd.PersistentFlags().StringVar(&loginOptions.Username, "username", "", "Optionally specify username. If not provided will be propted from via stdin")
	loginCmd.PersistentFlags().StringVar(&loginOptions.Password, "password", "", "Optionally specify password. if not provided will be propted from via stdin")
}
