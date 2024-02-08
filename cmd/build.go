package cmd

import (
	"github.com/bruttazz/devc/internal/builder"
	"github.com/bruttazz/devc/internal/configs"
	"github.com/spf13/cobra"
)

var buildOptions configs.BuildCmdOptions

var buildCmd = &cobra.Command{
	Use:   "build [options] [flags] env-name",
	Short: "build a devc env from a Dockerfile or Containerfile",
	Long: `build an environment from a Dockerfile or a Containerfile.
The file is intented to be available from the current working directoy. 
Otherwise specify the location of the file with "--file" flag.
If no specific arguments are provided, will use the current working directory as the 
build context and look for a Containerfile. The build fails if no
Containerfile nor Dockerfile is present.
`,
	Run: func(cmd *cobra.Command, args []string) {
		builder.Builder(&buildOptions, args)
	},
}

func init() {
	buildCmd.PersistentFlags().StringVarP(&buildOptions.Containerfile, "file", "f", "", "Explicitly point to a container file (can also provide valid urls)")
	buildCmd.PersistentFlags().StringVarP(&buildOptions.Context, "context", "c", ".", "Specify the build context to be used")
	buildCmd.PersistentFlags().BoolVar(&buildOptions.KeepCache, "keep-cache", false, "Keep build cache after successfull build of devc environement")
	buildCmd.PersistentFlags().StringArrayVar(&buildOptions.BuildArgs, "build-arg", []string{}, "`argument=value` to supply to the builder")
}
