package configs

// basic config options
var Config = &ConfigStruct{
	EnvSettings: EnvSettingsStruct{
		BuildDir: "builds",
		RootDir:  "root",
		DevcBin:  "bin",
	},
	CacheDirSettings: CacheDirSettingsStruct{
		Proot:         "proot",
		Buildah:       "from-builds",
		StorageDriver: "vfs",
	},
}
