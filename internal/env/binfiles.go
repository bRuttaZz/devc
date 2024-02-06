package env

import (
	"fmt"
	"path/filepath"

	"github.com/bruttazz/devc/internal/configs"
	"github.com/bruttazz/devc/internal/utils"
)

const activateString string = `
_DIR_NAME=%v
# deactivate other venv if exists
type deactivate &>/dev/null && deactivate

# setup for proot
_OLD_PATH=$PATH
_OLD_PS1=$PS1

_OLD_DEVC_WRKDIR=$DEVC_WRKDIR

if [ -z "${DEVC_WRKDIR}" ]
then
	DEVC_WRKDIR=/home/%v/devc
fi;

export PATH="$PATH:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"          
export PS1="(devc) $PS1"  

"$_DIR_NAME/proot" \
    -r "$_DIR_NAME/../root" \
    -b ".:$DEVC_WRKDIR" \
    -b "$SHELL:/bin/sh !" \
    -w "$DEVC_WRKDIR" \
    -0 \
    -b /dev \
    -b /proc \
    -b /sys \
    "/bin/sh"

# retaining the initial stage
export PATH=$_OLD_PATH
export PS1=$_OLD_PS1

DEVC_WRKDIR=$_OLD_DEVC_WRKDIR

unset _SELF_FILE_NAME
unset _DIR_NAME
unset _OLD_PATH
unset _OLD_PS1
unset _OLD_DEVC_WRKDIR
`

const deactivateString string = "#!/bin/sh\nkill -9 $PPID"

const nameServeString string = "nameserver 8.8.8.8\nnameserver 8.8.4.4"

// Add modifications to the container root
func finishUpRootBin(envPath string) (err error) {
	basePath := filepath.Join(envPath, configs.Config.EnvSettings.RootDir)

	// setup /etc/resolv.conf for assured internet access
	err = utils.WriteTextToFile(
		filepath.Join(basePath, "etc", "resolv.conf"),
		nameServeString,
	)
	if err != nil {
		return
	}

	// setup deactivate file to bin/deactivate
	err = utils.WriteTextToFile(
		filepath.Join(basePath, "bin", "deactivate"),
		deactivateString,
	)
	if err == nil {
		err = utils.MakeExecutable(filepath.Join(
			basePath,
			"bin",
			"deactivate",
		))
	}

	return
}

// Setup activate script to the devc bin
func setupActivateScript(envPath string) (err error) {
	activateScriptPath := filepath.Join(
		envPath,
		configs.Config.EnvSettings.DevcBin,
		"activate",
	)
	err = utils.WriteTextToFile(activateScriptPath, fmt.Sprintf(
		activateString,
		filepath.Join(envPath, configs.Config.EnvSettings.DevcBin),
		utils.CreateRandomString(),
	))
	if err != nil {
		err = utils.MakeExecutable(activateScriptPath)
	}
	return
}
