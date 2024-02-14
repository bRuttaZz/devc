package configs

// basic config options
var Config = &ConfigStruct{
	EnvSettings: EnvSettingsStruct{
		BuildDir: "builds",
		RootDir:  "root",
		DevcBin:  "bin",
	},
	CacheDirSettings: CacheDirSettingsStruct{
		ProotCache:    "proot",
		BuildahCache:  "from-builds",
		LoginAuthFile: "auth.json",
		StorageDriver: "vfs",
		CommonBuildCache: "build-states",
	},
}
