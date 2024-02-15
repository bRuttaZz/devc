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
type deactivate &>/dev/null && {
    echo "[WARNING] deactivating the existing env (it seems like an env is already present).";
    deactivate ;
}
# setup for proot
_OLD_PATH=$PATH
_OLD_PS1=$PS1
_OLD_SHELL=$SHELL

_OLD_DEVC_WRKDIR=$DEVC_WRKDIR

if [ -z "${DEVC_WRKDIR}" ]
then
	DEVC_WRKDIR=/home/%v/devc
fi;

# finding and fallbacking the default shell
if [ -f "$_DIR_NAME/../root/etc/passwd" ]
then 
    export SHELL=$(awk -F: -v user="root" '$1 == user {print $NF}' "$_DIR_NAME/../root/etc/passwd")
fi;
if [ -z "${SHELL}" ]
then
    export SHELL=/bin/sh
fi;

export PATH="$PATH:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"          
export PS1="(devc) $PS1"  
echo "export PS1=\"${PS1}\"" > "$_DIR_NAME/.rc";

"$_DIR_NAME/proot" \
    -r "$_DIR_NAME/../root" \
    -b ".:$DEVC_WRKDIR" \
	-b "$_DIR_NAME/deactivate:/bin/deactivate" \
    -b "$_DIR_NAME/.rc:/home/$USER/.$(basename $SHELL)rc" \
    -w "$DEVC_WRKDIR" \
    -0 \
	-b /dev \
    -b /proc \
    -b /sys \
    -b /tmp \
    -b /etc/host.conf \
    -b /etc/hosts \
	-b /etc/hostname \
    -b /etc/nsswitch.conf \
    -b /etc/resolv.conf \
    "$SHELL" ;

rm "$_DIR_NAME/.rc";
# retaining the initial stage
export PATH=$_OLD_PATH
export PS1=$_OLD_PS1
export SHELL=$_OLD_SHELL

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

	// setup /etc/resolv.conf for assured internet access (not required if mounting the file system)
	err = utils.WriteTextToFile(
		filepath.Join(basePath, "etc", "resolv.conf"),
		nameServeString,
	)
	return
}

// Setup activate script to the devc bin
func setupActivateScript(envPath string) (err error) {
	scriptPath := filepath.Join(
		envPath,
		configs.Config.EnvSettings.DevcBin,
		"activate",
	)
	err = utils.WriteTextToFile(scriptPath, fmt.Sprintf(
		activateString,
		filepath.Join(envPath, configs.Config.EnvSettings.DevcBin),
		utils.CreateRandomString(),
	))
	if err != nil {
		return
	}
	scriptPath = filepath.Join(
		envPath,
		configs.Config.EnvSettings.DevcBin,
		"deactivate",
	)
	err = utils.WriteTextToFile(
		scriptPath,
		deactivateString,
	)
	if err != nil {
		return
	}
	err = utils.MakeExecutable(scriptPath)
	return
}
