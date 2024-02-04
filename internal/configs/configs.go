package configs

import (
	"os/user"
	"path/filepath"

	"github.com/gobuffalo/packr"
	"gopkg.in/yaml.v2"
)

var Config ConfigStruct

func LoadConfig() {
	dat, err := packr.NewBox("../../configs").FindString("general.yml")
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
	Config.EnvSettings = EnvSettingsStruct{
		BuildDir: "builds",
		RootDir:  "root",
		DevcBin:  "bin",
	}
	Config.CacheDirSettings = CacheDirSettingsStruct{
		Proot:   "proot",
		Buildah: "from-builds",
	}
}
