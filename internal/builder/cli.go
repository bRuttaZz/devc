package builder

// the initial buildah imageBuilder based approach got some conflicts to resolve

import (
	"fmt"
	"path/filepath"

	"github.com/bruttazz/devc/internal/configs"
	"github.com/bruttazz/devc/internal/env"
)

func Builder(opts *configs.BuildCmdOptions, args []string) {
	if len(args) != 1 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}

	abs, err := filepath.Abs("")
	if err != nil {
		panic("[build error] error resolve path : " + err.Error())
	}

	var baseOptions = append(getGlobalBuildahOptions(abs, args[0]), getBuildOptions(abs, args[0])...)
	if len(opts.Containerfile) > 0 {
		baseOptions = append(baseOptions, "--file")
		baseOptions = append(baseOptions, filepath.Join(abs, opts.Containerfile))
	}
	baseOptions = append(baseOptions, filepath.Join(abs, opts.Context))

	// fmt.Println("options", buildCmd.Path, buildCmd.Args)
	err = runCommand(configs.Config.Buildah.Path, baseOptions)
	if err != nil {
		panic("[builder error] : " + err.Error())
	}
	fmt.Println("[devc] container created")
	err = env.SetupEnv(filepath.Join(abs, args[0]))
	if err != nil {
		panic("[setup error] : " + err.Error())
	}
	fmt.Printf("[devc] env created : %v\n", args[0])
}
