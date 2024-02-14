package builder

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bruttazz/devc/internal/configs"
	"github.com/bruttazz/devc/internal/utils"
)

const rootfsExtracterScript string = `#!/bin/sh 

set -e;

ctrName=$( __BUILDAH_BASE_EXEC from __IMAGE_NAME )
mountName=$( __BUILDAH_BASE_EXEC mount $ctrName )
mkdir -p $(dirname __OUT_ROOT)
cp -r $mountName __OUT_ROOT
chmod -R u+rwx __OUT_ROOT
__BUILDAH_BASE_EXEC umount $ctrName 
__BUILDAH_BASE_EXEC rm $ctrName  
`

func exportImageAsRootFs(imageName, outDirPath string) (err error) {
	var cmdLines = configs.Config.Buildah.Path + " " + strings.Join(getGlobalBuildahOptions(outDirPath), " ")
	cmdLines = strings.ReplaceAll(rootfsExtracterScript, "__BUILDAH_BASE_EXEC", cmdLines)
	cmdLines = strings.ReplaceAll(cmdLines, "__IMAGE_NAME", imageName)
	cmdLines = strings.ReplaceAll(cmdLines, "__OUT_ROOT", filepath.Join(outDirPath, configs.Config.EnvSettings.RootDir))

	// actual tempfile got issue on forking the defer command gets to do something
	os.MkdirAll(configs.Config.CacheDir, 0755)
	fileName := filepath.Join(configs.Config.CacheDir, "_temp_script_"+imageName+".sh")
	err = utils.WriteTextToFile(fileName, cmdLines)
	defer os.Remove(fileName)
	if err != nil {
		return
	}
	err = utils.MakeExecutable(fileName)
	if err != nil {
		return
	}
	err = runCommand(configs.Config.Buildah.Path, []string{
		"unshare",
		fileName,
	})
	return
}
