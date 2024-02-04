package configs

type EnvSettingsStruct struct {
	BuildDir string
	RootDir  string
	DevcBin  string
}

type CacheDirSettingsStruct struct {
	Proot   string
	Buildah string
}

type ConfigStruct struct {
	Proot struct {
		Url     string `yaml:"url"`
		Version string `yaml:"version"`
	} `yaml:"proot"`

	CacheDir         string  `yaml:"cache-dir"`
	CacheExpiryHrs   float64 `yaml:"cache-expiry-hrs"`
	HomeDir          string
	UserName         string
	EnvSettings      EnvSettingsStruct
	CacheDirSettings CacheDirSettingsStruct
}
