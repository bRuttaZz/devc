package internal

import (
	"github.com/gobuffalo/packr"
	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	Proot struct {
		Url     string `yaml:"url"`
		Version string `yaml:"version"`
	} `yaml:"proot"`

	CacheDir string `yaml:"cache-dir"`
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
}
