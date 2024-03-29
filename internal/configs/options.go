package configs

type RootOptions struct {
	Version bool
}

type ActivateCmdOptions struct {
	Workdir string
}

type BuildCmdOptions struct {
	Containerfile string
	Context       string
	KeepCache     bool
	BuildArgs     []string
}

type ImagesCmdOptions struct {
}

type LoginCmdOptions struct {
	Username string
	Password string
}

type LogoutCmdOptions struct {
}

type PruneCmdOptions struct {
	Wipe bool
}

type PullCmdOptions struct {
	NoCaching bool
}

type RmiCmdOptions struct {
}
