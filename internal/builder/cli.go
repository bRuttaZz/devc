package builder

// the initial buildah imageBuilder based approach got some conflicts to resolve

import (
	"fmt"
	"os"
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
	_, err = os.Stat(filepath.Join(abs, args[0]))
	if err == nil {
		panic("[ERROR] directory already exists! '" + args[0] + "'")
	}
	var globalOptions = getGlobalBuildahOptions(args[0])
	imgName, buildOptions, err := getBuildOptions()
	if err != nil {
		panic("[devc error] getting builder options : " + err.Error())
	}
	buildOptions = append(globalOptions, buildOptions...)
	for i := 0; i < len(opts.BuildArgs); i++ {
		buildOptions = append(buildOptions, []string{"--build-arg", opts.BuildArgs[i]}...)
	}

	if len(opts.Containerfile) > 0 {
		buildOptions = append(buildOptions, "--file")
		buildOptions = append(buildOptions, filepath.Join(abs, opts.Containerfile))
	}
	buildOptions = append(buildOptions, filepath.Join(abs, opts.Context))

	fmt.Println("[devc] building container..")
	err = runCommand(configs.Config.Buildah.Path, buildOptions)
	if err != nil {
		panic("[builder error] stage 1 : " + err.Error())
	}
	fmt.Println("[devc] image created..")
	fmt.Println("[devc] creating devc env ")

	err = exportImageAsRootFs(imgName, filepath.Join(abs, args[0]))
	if err != nil {
		panic("[devc] mount error : " + err.Error())
	}

	err = env.SetupEnv(filepath.Join(abs, args[0]))
	if err != nil {
		panic("[setup error] : " + err.Error())
	}
	fmt.Printf("[devc] env created : %v\n", args[0])
	if !opts.KeepCache {
		defer func() {
			if v := recover(); v != nil {
				panic(fmt.Sprint("[finishup error] : ", v))
			}
		}()
		Rmi(&configs.RmiCmdOptions{}, []string{imgName})
	}
	fmt.Printf("\n[devc] tada! the env is all yours : [%v]\n", args[0])
}

// activate the venv
func Activate(opts *configs.ActivateCmdOptions, args []string) {
	if len(args) != 1 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}

	if len(opts.Workdir) > 0 {
		os.Setenv("DEVC_WRKDIR", opts.Workdir)
	}
	err := runCommand(os.Getenv("SHELL"), []string{
		"-c",
		fmt.Sprintf("PS1=\"[\\u@\\h \\W] \\$ \" source %v",
			filepath.Join(
				args[0],
				configs.Config.EnvSettings.DevcBin,
				"activate",
			),
		),
	})
	if err != nil {
		panic("[devc activate error] : " + err.Error())
	}
}

// pull images and create env
func Pull(opts *configs.PullCmdOptions, args []string) {
	if len(args) != 2 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}

	abs, err := filepath.Abs("")
	if err != nil {
		panic("[pull error] error resolve path : " + err.Error())
	}
	_, err = os.Stat(filepath.Join(abs, args[1]))
	if err == nil {
		panic("[ERROR] directory already exists! '" + args[1] + "'")
	}
	err = exportImageAsRootFs(args[0], filepath.Join(abs, args[1]))
	if err != nil {
		panic("[devc] mount error : " + err.Error())
	}

	err = env.SetupEnv(filepath.Join(abs, args[1]))
	if err != nil {
		panic("[setup error] : " + err.Error())
	}
	fmt.Printf("[devc] env created : %v\n", args[1])
	if opts.NoCaching {
		defer func() {
			if v := recover(); v != nil {
				panic(fmt.Sprint("[finishup error] : ", v))
			}
		}()
		Rmi(&configs.RmiCmdOptions{}, []string{args[0]})
	}
	fmt.Printf("\n[devc] tada! the env is all yours : [%v]\n", args[1])

}

// login to a registry
func Login(opts *configs.LoginCmdOptions, args []string) {
	if len(args) != 1 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}
	var baseOptions = getGlobalBuildahOptions("")
	baseOptions = append(baseOptions, "login")
	authFileOpts, err := getAuthFileOptions()
	if err != nil {
		panic("[devc error] : " + err.Error())
	}
	baseOptions = append(baseOptions, authFileOpts...)
	if len(opts.Username) > 0 {
		baseOptions = append(baseOptions, []string{
			"--username",
			opts.Username,
		}...)
	}
	if len(opts.Password) > 0 {
		baseOptions = append(baseOptions, []string{
			"--password",
			opts.Password,
		}...)
	}
	baseOptions = append(baseOptions, args[0])
	// fmt.Println("options", buildCmd.Path, buildCmd.Args)
	err = runCommand(configs.Config.Buildah.Path, baseOptions)
	if err != nil {
		panic("[login error] : " + err.Error())
	}
	fmt.Printf("[devc login] successfully logged into \"%v\"\n", args[0])
}

// logout from registry
func Logout(opts *configs.LogoutCmdOptions, args []string) {
	if len(args) != 1 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}
	var baseOptions = getGlobalBuildahOptions("")
	baseOptions = append(baseOptions, "logout")
	authFileOpts, err := getAuthFileOptions()
	if err != nil {
		panic("[devc error] : " + err.Error())
	}
	baseOptions = append(baseOptions, authFileOpts...)
	baseOptions = append(baseOptions, args[0])
	// fmt.Println("options", buildCmd.Path, buildCmd.Args)
	err = runCommand(configs.Config.Buildah.Path, baseOptions)
	if err != nil {
		panic("[logout error] : " + err.Error())
	}
	fmt.Printf("[devc logout] successfully out from \"%v\"\n", args[0])
}

// list all cached images
func Images(opts *configs.ImagesCmdOptions, args []string) {
	if len(args) != 0 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}

	var baseOptions = getGlobalBuildahOptions("")
	baseOptions = append(baseOptions, "images")
	// fmt.Println("options", buildCmd.Path, buildCmd.Args)
	err := runCommand(configs.Config.Buildah.Path, baseOptions)
	if err != nil {
		panic("[image list error] : " + err.Error())
	}
}

// Prune home cache dir
func Prune(opts *configs.PruneCmdOptions, args []string) {
	if len(args) != 0 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}
	err := clearAllImageCache()
	if err != nil {
		panic("[devc prune error] : " + err.Error())
	}
	if opts.Wipe {
		os.Remove(filepath.Join(configs.Config.CacheDir, configs.Config.CacheDirSettings.ProotCache))
		os.Remove(filepath.Join(configs.Config.CacheDir, configs.Config.CacheDirSettings.LoginAuthFile))
	}
	fmt.Println("[devc] system prune complete!")
}

// Remove a cached image (cache is only applicable for devc pull command)
func Rmi(opts *configs.RmiCmdOptions, args []string) {
	if len(args) != 1 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}
	options := getGlobalBuildahOptions("")
	options = append(options, "rmi")
	options = append(options, args[0])

	err := runCommand(configs.Config.Buildah.Path, options)
	if err != nil {
		panic("[devc rmi error] : " + err.Error())
	}
	fmt.Printf("[devc] successfully removed image (%v)! \n", args[0])
}
