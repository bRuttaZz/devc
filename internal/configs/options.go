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
}

type PullCmdOptions struct {
}

type ImagesCmdOptions struct {
}

type PruneCmdOptions struct {
	Wipe bool
}

type RmCmdOptions struct {
}

type RmiCmdOptions struct {
}
