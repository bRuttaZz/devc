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
	var baseOptions = append(getGlobalBuildahOptions(abs, args[0]), getBuildOptions(abs, args[0])...)
	authFileOpts, err := getAuthFileOptions()
	if err != nil {
		panic("[devc error] : " + err.Error())
	}
	baseOptions = append(baseOptions, authFileOpts...)
	for i := 0; i < len(opts.BuildArgs); i++ {
		baseOptions = append(baseOptions, []string{"--build-arg", opts.BuildArgs[i]}...)
	}

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
	if !opts.KeepCache {
		err = clearBuildCache(args[0])
		if err != nil {
			panic("[finishup error] : " + err.Error())
		}
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
		fmt.Sprintf("PS1=\"[\\u@\\h \\W]\\$\" source %v",
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
	file, err := os.CreateTemp("", "pullcmd")
	if err != nil {
		panic("[pull error] : " + err.Error())
	}
	defer file.Close()
	defer os.Remove(file.Name())
	_, err = file.Write([]byte(fmt.Sprintf("FROM %v", args[0])))
	if err != nil {
		panic("[pull error] : " + err.Error())
	}

	abs, err := filepath.Abs("")
	if err != nil {
		panic("[pull error] error resolve path : " + err.Error())
	}

	var baseOptions = append(getGlobalBuildahOptions("", ""), getBuildOptions(abs, args[1])...)
	authFileOpts, err := getAuthFileOptions()
	if err != nil {
		panic("[devc error] : " + err.Error())
	}
	baseOptions = append(baseOptions, authFileOpts...)

	baseOptions = append(baseOptions, "--file")
	baseOptions = append(baseOptions, file.Name())
	baseOptions = append(baseOptions, ".")

	// fmt.Println("options", buildCmd.Path, buildCmd.Args)
	err = runCommand(configs.Config.Buildah.Path, baseOptions)
	if err != nil {
		panic("[pull error] : " + err.Error())
	}
	fmt.Println("[devc] container created..")
	err = env.SetupEnv(filepath.Join(abs, args[1]))
	if err != nil {
		panic("[setup error] : " + err.Error())
	}
	if opts.NoCaching {
		baseOptions = getGlobalBuildahOptions("", "")
		baseOptions = append(baseOptions, "rmi")
		baseOptions = append(baseOptions, args[0])

		err := runCommand(configs.Config.Buildah.Path, baseOptions)
		if err != nil {
			panic("[devc cleanup error] : " + err.Error())
		}
		fmt.Println("[devc] pull cache removed..")
	}
	fmt.Printf("\n[devc] tada! the env is all yours : [%v]\n", args[1])
}

// login to a registry
func Login(opts *configs.LoginCmdOptions, args []string) {
	if len(args) != 1 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}
	var baseOptions = getGlobalBuildahOptions("", "")
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
	fmt.Printf("[devc login] successfully loged into \"%v\"\n", args[0])
}

// logout from registry
func Logout(opts *configs.LogoutCmdOptions, args []string) {
	if len(args) != 1 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}
	var baseOptions = getGlobalBuildahOptions("", "")
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

	var baseOptions = getGlobalBuildahOptions("", "")
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
	err := clearBuildCache("")
	if err != nil {
		panic("[devc prune error] : " + err.Error())
	}
	if opts.Wipe {
		os.Remove(filepath.Join(configs.Config.CacheDir, configs.Config.CacheDirSettings.ProotCache))
		os.Remove(filepath.Join(configs.Config.CacheDir, configs.Config.CacheDirSettings.LoginAuthFile))
	}
	fmt.Println("[devc] system prune complete!")
}

// Remove an existing devc env (simply deleting it will also work)
func Rm(opts *configs.RmCmdOptions, args []string) {
	if len(args) != 1 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}
	// check if it's a valid env
	_, err := os.Stat(filepath.Join(args[0], configs.Config.EnvSettings.DevcBin, "activate"))
	if err != nil {
		panic("[devc rm error] : seems not to be a devc environment")
	}
	err = clearBuildCache(args[0])
	if err != nil {
		panic("[devc rm error] : " + err.Error())
	}
	err = runCommand("rm", []string{"-rf", args[0]})
	if err != nil {
		panic("[devc rm error] : " + err.Error())
	}
	fmt.Println("[devc] env removed successfully!")
}

// Remove a cached image (cache is only applicable for devc pull command)
func Rmi(opts *configs.RmiCmdOptions, args []string) {
	if len(args) != 1 {
		panic("Invalid number of positional argument. Execute command with --help to get detailed usecase")
	}
	options := getGlobalBuildahOptions("", "")
	options = append(options, "rmi")
	options = append(options, args[0])

	err := runCommand(configs.Config.Buildah.Path, options)
	if err != nil {
		panic("[devc rmi error] : " + err.Error())
	}
	fmt.Printf("[devc] successfully removed image (%v)! \n", args[0])
}
