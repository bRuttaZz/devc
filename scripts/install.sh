#!/bin/sh
set -e

MAN="devC installation script.\n
\n
This script is meant to be a convenient way to install devc in linux systems\n
\n
The script:\n
\n
- Requires 'root' or 'sudo' privileges to run.\n
- Attempts to detect your Linux distribution and version and configure your
  package management system for you.\n
- Doesn't allow you to customize most installation parameters.\n
- Installs dependencies and recommendations without asking for confirmation.\n
- Installs the latest stable release (by default) of devc, buildah, and runc.\n
- Isn't designed to upgrade an existing devc installation. 
\n
Source code is available at https://github.com/bRuttaZz/devc/blob/main/scripts/install.sh\n
\n
\n
USAGE\n
=====
To install the latest stable versions of devc, and its dependencies: \n
1. download the script\n
  \t$ curl -fsSL https://raw.githubusercontent.com/bRuttaZz/devc/main/scripts/install.sh -o install-devc.sh\n
2. verify the script's content\n
  \t$ cat install-devc.sh\n
3. run the script with --dry-run to verify the steps it executes\n
  \t$ sh install-devc.sh --dry-run\n
4. run the script either as root, or using sudo to perform the installation.\n
  \t$ sudo sh install-devc.sh\n
\n
Command-line options\n
====================\n
--dry-run\n
\tUse to run a dry test for the installation script.\n

--bin-path=<path>\n
\tTo specify devc installation path. Defaults to /usr/local/bin .\n

--help | -h\n
\tTo get detailed usage details of this script.\n
"

DRY_RUN=${DRY_RUN:-}
BIN_PATH=/usr/local/bin
while [ $# -gt 0 ]; do
	case "$1" in
		--dry-run)
			DRY_RUN=1
			;;
		--bin-path)
			BIN_PATH=$2
			;;
		--help|-h)
			echo -e $MAN
			exit 0
			;;
		--*)
			echo "Illegal option $1"
			exit 1
			;;
	esac
	shift $(( $# > 0 ? 1 : 0 ))
done

command_exists() {
	command -v "$@" > /dev/null 2>&1
}

is_dry_run() {
	if [ -z "$DRY_RUN" ]; then
		return 1
	else
		return 0
	fi
}

is_wsl() {
	case "$(uname -r)" in
	*microsoft* ) true ;; # WSL 2
	*Microsoft* ) true ;; # WSL 1
	* ) false;;
	esac
}

is_darwin() {
	case "$(uname -s)" in
	*darwin* ) true ;;
	*Darwin* ) true ;;
	* ) false;;
	esac
}

deprecation_notice() {
	distro=$1
	distro_version=$2
	echo
	printf "# \033[91;1mDEPRECATION WARNING\033[0m\n"
	printf "#    This Linux distribution (\033[1m%s %s\033[0m) reached end-of-life and is no longer supported by this script.\n" "$distro" "$distro_version"
	echo   "#    No updates or security fixes will be released for this distribution, and users are recommended"
	echo   "#    to upgrade to a currently maintained version of $distro."
	echo   "#"
	printf   "# Press \033[1mCtrl+C\033[0m now to abort this script, or wait for the installation to continue."
	echo
	if ! is_dry_run; then
		sleep 10
	fi
}

get_distribution() {
	lsb_dist=""
	# Every system that we officially support has /etc/os-release
	if [ -r /etc/os-release ]; then
		lsb_dist="$(. /etc/os-release && echo "$ID")"
	fi
	# Returning an empty string here should be alright since the
	# case statements don't act unless you provide an actual value
	echo "$lsb_dist"
}

# Check if this is a forked Linux distro
check_forked() {

	# Check for lsb_release command existence, it usually exists in forked distros
	if command_exists lsb_release; then
		# Check if the `-u` option is supported
		set +e
		lsb_release -a -u > /dev/null 2>&1
		lsb_release_exit_code=$?
		set -e

		# Check if the command has exited successfully, it means we're in a forked distro
		if [ "$lsb_release_exit_code" = "0" ]; then
			# Print info about current distro
			cat <<-EOF
			You're using '$lsb_dist' version '$dist_version'.
			EOF

			# Get the upstream release info
			lsb_dist=$(lsb_release -a -u 2>&1 | tr '[:upper:]' '[:lower:]' | grep -E 'id' | cut -d ':' -f 2 | tr -d '[:space:]')
			dist_version=$(lsb_release -a -u 2>&1 | tr '[:upper:]' '[:lower:]' | grep -E 'codename' | cut -d ':' -f 2 | tr -d '[:space:]')

			# Print info about upstream distro
			cat <<-EOF
			Upstream release is '$lsb_dist' version '$dist_version'.
			EOF
		else
			if [ -r /etc/debian_version ] && [ "$lsb_dist" != "ubuntu" ] && [ "$lsb_dist" != "raspbian" ]; then
				if [ "$lsb_dist" = "osmc" ]; then
					# OSMC runs Raspbian
					lsb_dist=raspbian
				else
					# We're Debian and don't even know it!
					lsb_dist=debian
				fi
				dist_version="$(sed 's/\/.*//' /etc/debian_version | sed 's/\..*//')"
				case "$dist_version" in
					12)
						dist_version="bookworm"
					;;
					11)
						dist_version="bullseye"
					;;
					10)
						dist_version="buster"
					;;
					9)
						dist_version="stretch"
					;;
					8)
						dist_version="jessie"
					;;
				esac
			fi
		fi
	fi
}

# require to install depends and exit
require_installation() {
	echo -e "#\n# ERROR [devc] : dependecy installation failed for '$1'.
	# Kindly install the dependecies and try again with the script!"
	exit 1
}

get_shc() {
	user="$(id -un 2>/dev/null || true)"
	sh_c='sh -c'
	if [ "$user" != 'root' ]; then
		if command_exists sudo; then
			sh_c='sudo -E sh -c'
		elif command_exists su; then
			sh_c='su -c'
		else
			cat >&2 <<-'EOF'
			#
			# [ERROR] this installer needs the ability to run commands as root.
			# We are unable to find either "sudo" or "su" available to make this happen.
			EOF
			exit 1
		fi
	fi

	if is_dry_run; then
		sh_c="echo"
	fi
	printf $sh_c
}

get_arch() {
	case $(uname -m) in 
		x86_64)
			printf "amd64"
			;;
		i686)
			printf "386"
			;;
		aarch64)
			printf "arm64"
			;;
		arm)
			printf "arm"
			;;
		*)
			printf "amd64"
			;;
	esac
}

install_dependencies() {

	if command_exists devc; then
		cat >&2 <<-'EOF'
			# Warning: the "devc" command appears to already exist on this system.
#
#           Trying to REINSTALL devc.
			# You may press Ctrl+C now to abort this script.
		EOF
		if ! is_dry_run; then
			( set -x; sleep 10 )
		fi;
	fi

	sh_c=$( get_shc )
    # perform some very rudimentary platform detection
	lsb_dist=$( get_distribution )
	lsb_dist="$(echo "$lsb_dist" | tr '[:upper:]' '[:lower:]')"

	if is_wsl; then
		echo
		echo "# WSL DETECTED: Think like you are trying to install devc inside windows."
		echo "# This setup is not properly tested out. Forward with caution! "
        echo "# I mean the installation may break :)"
		echo
		cat >&2 <<-'EOF'
			
			# You may press Ctrl+C now to abort this script.
		EOF
		if ! is_dry_run; then
			( set -x; sleep 20 )
		fi
	fi


    ############################################
    # installing requirements
    if command_exists "buildah" ; then
        echo "# [requirement] devc requirements already satisfied!"
    else
        echo "# [requirement] installing requirements!"

        # get exact distros
        case "$lsb_dist" in

            ubuntu)
                if command_exists lsb_release; then
                    dist_version="$(lsb_release --codename | cut -f2)"
                fi
                if [ -z "$dist_version" ] && [ -r /etc/lsb-release ]; then
                    dist_version="$(. /etc/lsb-release && echo "$DISTRIB_CODENAME")"
                fi
            ;;

            debian|raspbian)
                dist_version="$(sed 's/\/.*//' /etc/debian_version | sed 's/\..*//')"
                case "$dist_version" in
                    12)
                        dist_version="bookworm"
                    ;;
                    11)
                        dist_version="bullseye"
                    ;;
                    10)
                        dist_version="buster"
                    ;;
                    9)
                        dist_version="stretch"
                    ;;
                    8)
                        dist_version="jessie"
                    ;;
                esac
            ;;

            centos|rhel|sles)
                if [ -z "$dist_version" ] && [ -r /etc/os-release ]; then
                    dist_version="$(. /etc/os-release && echo "$VERSION_ID")"
                fi
            ;;

            *)
                if command_exists lsb_release; then
                    dist_version="$(lsb_release --release | cut -f2)"
                fi
                if [ -z "$dist_version" ] && [ -r /etc/os-release ]; then
                    dist_version="$(. /etc/os-release && echo "$VERSION_ID")"
                fi
            ;;

        esac

        # Check if this is a forked Linux distro
        check_forked

        # Print deprecation warnings for distro versions that recently reached EOL,
        # but may still be commonly used (especially LTS versions).
        case "$lsb_dist.$dist_version" in
            debian.stretch|debian.jessie)
                deprecation_notice "$lsb_dist" "$dist_version"
                ;;
            raspbian.stretch|raspbian.jessie)
                deprecation_notice "$lsb_dist" "$dist_version"
                ;;
            ubuntu.xenial|ubuntu.trusty)
                deprecation_notice "$lsb_dist" "$dist_version"
                ;;
            ubuntu.impish|ubuntu.hirsute|ubuntu.groovy|ubuntu.eoan|ubuntu.disco|ubuntu.cosmic)
                deprecation_notice "$lsb_dist" "$dist_version"
                ;;
            fedora.*)
                if [ "$dist_version" -lt 36 ]; then
                    deprecation_notice "$lsb_dist" "$dist_version"
                fi
                ;;
        esac

	    # Run setup for each distro accordingly
	    case "$lsb_dist" in
		    ubuntu|debian|raspbian)
				{
					$sh_c "apt-get update" &&
					$sh_c "apt-get install -y buildah"
				} || {
					require_installation buildah
				}
			    ;;
            centos|rhel)
				{
	                $sh_c "yum install -y buildah"
				} || {
					require_installation buildah
				}
                ;;
    		fedora)
				{
	                $sh_c "dnf install -y buildah"
				} || {
					require_installation buildah
				}
		    	;;
            arch)
                {
					$sh_c "pacman -Sy buildah"
				} || {
					require_installation buildah
				}
                ;;
			*)
				if command_exists "emerge"; then
					{
						$sh_c "emerge app-containers/buildah"
					} || {
						require_installation buildah
					}
				elif command_exists "zypper"; then
					{
						$sh_c "zypper install buildah"
					} || {
						require_installation buildah
					}
				elif [ -z "$lsb_dist" ]; then
					if is_darwin; then
						echo "#"
						echo "# ERROR: Unsupported operating system 'macOS'"
						echo "# Sorry FOLKS :)"
						echo 
						exit 1
					fi
				else 
					echo "#"
					echo "# [ERROR]: Unsupported distribution '$lsb_dist'"
					echo
					exit 1
				fi
				;;
	    esac
    fi
    ############################################
}

install_devc() {
	set -e
	# download file
	echo "# [devc] downloading latest binary.."
	url=$( curl -s https://api.github.com/repos/bruttazz/devc/releases/latest \
		| grep "browser_download_url" \
		| grep -i "devc.*$(get_arch)" \
		| cut -d : -f 2,3 \
		| tr -d \" )
	if [ -z $url ]; then
		echo -e "#\n# [ERROR] unable to find devc binary compiled for your system from the latest release.
# Consider installing from source :)"
		exit 1
	fi
	if is_dry_run ; then
		executer="echo"
	fi
	$executer wget -O devc "$url";
	$executer "chmod a+x ./devc" ;
	$(get_shc) "cp ./devc $BIN_PATH/devc" ;
	echo -e "# [devc] Installation successfull!
#\tTry executing 'devc --version' to verify the installation or 'devc --help' for more info."
}

do_install() {
	echo "# [ Executing devc install script ]"
	install_dependencies
	install_devc
}

# wrapped up in a function so that we have some protection against only getting
# half the file during "curl | sh"
do_install

