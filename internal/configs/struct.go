package configs

import (
	"os/user"
	"path/filepath"

	"github.com/bruttazz/devc/config"
	"gopkg.in/yaml.v2"
)

type EnvSettingsStruct struct {
	BuildDir string
	RootDir  string
	DevcBin  string
}

type CacheDirSettingsStruct struct {
	ProotCache    string
	BuildahCache  string // for dumping image data
	LoginAuthFile string

	StorageDriver    string
	CommonBuildCache string // for dumping state dat
}

type ConfigStruct struct {
	Version string
	Proot   struct {
		Url     string `yaml:"url"`
		Version string `yaml:"version"`
	} `yaml:"proot"`

	CacheDir         string  `yaml:"cache-dir"`
	CacheExpiryHrs   float64 `yaml:"cache-expiry-hrs"`
	HomeDir          string
	UserName         string
	EnvSettings      EnvSettingsStruct
	CacheDirSettings CacheDirSettingsStruct

	Buildah struct {
		Path string `yaml:"path"`
	} `yaml:"buildah"`
}

func LoadConfig() {
	if err := yaml.Unmarshal(config.General, &Config); err != nil {
		panic("error parsing devc config : " + err.Error())
	}

	usr, err := user.Current()
	if err != nil {
		panic("error resolving current user : " + err.Error())
	}
	Config.HomeDir = usr.HomeDir
	Config.UserName = usr.Username
	Config.CacheDir = filepath.Join(Config.HomeDir, Config.CacheDir)
}
