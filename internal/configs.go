package internal

import (
	"os/user"
	"path/filepath"

	"github.com/gobuffalo/packr"
	"gopkg.in/yaml.v2"
)

type EnvSettings struct {
	BuildDir string
	RootDir  string
	DevcBin  string
}

type ConfigStruct struct {
	Proot struct {
		Url     string `yaml:"url"`
		Version string `yaml:"version"`
	} `yaml:"proot"`

	CacheDir       string  `yaml:"cache-dir"`
	CacheExpiryHrs float64 `yaml:"cache-expiry-hrs"`
	HomeDir        string
	UserName       string
	EnvSettings    EnvSettings
}

var Config ConfigStruct

func LoadConfig() {
	dat, err := packr.NewBox("../configs").FindString("general.yml")
	if err != nil {
		panic("error loading devc config : " + err.Error())
	}
	if err := yaml.Unmarshal([]byte(dat), &Config); err != nil {
		panic("error parsing devc config : " + err.Error())
	}

	usr, err := user.Current()
	if err != nil {
		panic("error resolving current user : " + err.Error())
	}
	Config.HomeDir = usr.HomeDir
	Config.UserName = usr.Username
	Config.CacheDir = filepath.Join(Config.HomeDir, Config.CacheDir)
	Config.EnvSettings = EnvSettings{
		BuildDir: "builds",
		RootDir:  "root",
		DevcBin:  "bin",
	}
}
