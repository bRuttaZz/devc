package configs

type RootOptions struct {
	Verbose bool
}

type BuildCmdOptions struct {
	Containerfile string
	Context       string
	KeepCache     bool
}
