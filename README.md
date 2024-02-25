# devC

<a href="./assets/LICENSE.md" title="Logo by Dev G"><img src="./assets/tad.png" width=150 alt="Logo by Dev G"/></a>

**Containers For Developers (Container as a Folder)**

![Workflow status](https://github.com/bruttazz/devc/actions/workflows/codeql.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/bruttazz/devc)](https://goreportcard.com/report/github.com/bruttazz/devc)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://github.com/bruttazz/devc/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/bruttazz/devc.svg)](https://pkg.go.dev/github.com/bruttazz/devc)
[![Release](https://img.shields.io/github/release/bruttazz/devc.svg?style=flat-square)](https://github.com/bruttazz/devc/releases/latest)

---

<!-- <details>  -->
<!-- <summary><b>Quick Links</b></summary> -->

## TL;DR 
using **devc** one can create a virtualenv (similar to python venv, as a directory) from a Dockerfile / Containerfile, or directly from a docker image resides on a container regisrty of your choice. Then activate the env, edit your code and run it inside the (jailed) container-like environment. You can install new things, code your app using any containers available, all in user space. (No deamons included :). <br>Jump to [usage section](#usage)


## Quick Links

1. [Intro](#intro)
2. [Usage](#usage)
3. [Installation](#installation)
4. [Uninstalling](#uninstalling)
5. [Known Issues](#known-issues)
6. [Credits](#credits)

<!-- </details> -->

Detailed (I mean semi-detailed) [CLI usage manual](./docs/man.md), and [release notes](./docs/CHANGELOG.md).

Found an issue? Let's discus,<br>
Matrix group : [#devc:matrix.org](https://matrix.to/#/!nEmTMcQUkCipApdYVE:matrix.org?via=matrix.org) <br>
Github : [Discussions](https://github.com/bRuttaZz/devc/discussions), [Issues](https://github.com/bRuttaZz/devc/issues)

Found an issue? Let's discus,<br>
Matrix group : [#devc:matrix.org](https://matrix.to/#/!nEmTMcQUkCipApdYVE:matrix.org?via=matrix.org) <br>
Github : [Discussions](https://github.com/bRuttaZz/devc/discussions), [Issues](https://github.com/bRuttaZz/devc/issues)

## Intro

Have you ever utilized Docker or Podman in local application development, just for resolving the dependency issues? If so, have you encountered challenges in creating and maintaining the Docker environment for development? This may involve rebuilding your project with each Dockerfile change, mounting your local codebase to the container, dealing with permission overhead, and connecting your IDE to the container.

The good news is, if you're on Linux, you can give **devc** a try. It's a streamlined container tool designed for developers working on local development. It's compatible with contemporary container standards like Docker and OCI. **devc** facilitates you with a python-virtual-environment kind of way to interact with container environment. 
Basically, using devc one can create a **new venv similar to python-venv** from existing Dokckerfile/Containerfile or by specifying a docker or oci image at a registry (eg: docker-hub). 
After creation the developer can activate that env using `source venv/bin/activate` (yes, exactly like python venv). And (if everything went well) from the next instance your terminal prompt will be prefixed with `(devc)` and you will be inside the newly created container environment, together with all your files and directories from the current working directory. 

The venv may of different linux distribution with different setup, whatever be that, one can install any application they want to it, without affecting the host machine. After you are tired up by playing with it, simply type `deactivate` (yeah you got it, it's quite python-venv like), or `ctrl + d` keyboard combination to exit from the venv, and you will be on your normal original terminal. Want to enter again? type the same command `source venv/bin/activate`, all your previously made changes will be waiting over there. And at the end of the day. You can simply delete the venv directory. And that's how you going to manage the storage (just like in the good old days before the storage hungry container monsters).

**NB** : In case you missed it, **devc** is just a simple developer tool, it's meant to be used for local development, and it's a striped down version of normal container setups. 

## Yet Another Container Tool ?
Okay you may be thinking **devc** as a wrap around bunch of Docker or Podman workflows to present it similar to the way of python venv. In fact devc handles things in a different way. Docker, Podman or other similar container tools are optimized to work with deployment time, and they got a lot of process management and resource sharing features fine-tuned dedicatedly for that as well. But in devc we are taking only the sand-boxed rootfs and binding system (we simply omit all other features :) 

Following points may detail about the behavior of devc

1. Devc will be maintaining a standalone *rootfs*, so that the users will be able to install and run different programs without affecting the host system, other than where the current working directory is situated.
2. Unlike other container tools, Devc does not have any deamon processes of its own. Every command issued on devc environment will be run as in host namespace. Making it lightweight, and developer friendly.
3. Devc can run completely on user space. It's absolutely **not recommended to run devc with superuser privileges**
4. The container build cache and the intermediate images pulled while building a devc environment will be stored under the venv (although the cache directories will be wiped as soon as the container images are built). Deleting it will delete all other cache dependent upon the venv.
5. The current working directory on activating time will be mounted to the devc environment at a randomly created 'default' workdir of devc (All other file system will be in a sand-boxed environment). Due to this, one can seamlessly make changes in the files of current work directory and test it out using the interactive terminal of devc inside the activated container environment.

**About the Name** : As you may have guessed by now, the name "devc" is made of two words "developer" and "container". i.e., **containers for developers, but here developers got priority :)**

## Usage

After installation (BTW, you can find the detailed installation procedure over [here](#installation)), execute
```sh
devc --version
```
in your terminal to test if everything went well. For detailed usage manual refer [devc manual](./docs/man.md)

### 1. Creating a devc env
One can create a devc env either by pulling a docker/OCI container image from a container registry, or by building a new one from a Dockerfile or Containerfile.

#### 1.1 Creating a devc env from existing docker image
Execute 
```sh
devc pull <image-name> <env-name>
```
where `image-name` represents the name of the image to be pulled, it can also contain version tag and registry names.
`env-name` will be the name of your new devc environment. It's going to create a directory with the exact name in your current workspace.

eg: to create a simple python devc env, one can execute `devc pull alpine env` it will create a new directory called `env`, with docker image pulled from any of the default container registry configured on your system. If you want to specify the registry and version tag on pulling, execute something like `devc pull docker.io/library/python:3.12-alpine env`.

#### 1.2 Creating a devc env from Dockerfile/Containerfile
Navigate into the directory of having your Dockerfile / Containerfile and execute the following to build and create a new devc container env
```sh
devc build <env-name>
```

or if you want to specify a Dockerfile at different location or an http url of the build file, try out
```sh
devc build -f <Dockerfile/url> <env-name>
```

eg: executing the command `devc build env` will create a new devc container env named `env`, at the current working directory from a Containerfile / Dockerfile from the working directory, if exists.

### 2. Activating the devc env
After successful creation one can activate devc env just like a python virtualenv. execute
```sh
source <env-name>/bin/activate
``` 
If everything went well, your terminal prompt will be prefixed with `(devc)` (the behavior may change if the container's default shell is not supporting much modifications. In such scenarios it's recommended to go with `devc activate` command. Refer the [manual](./docs/man.md) for more info).

### 3. Deactivating the devc terminal
Again simple, as devc is not running any demon processes, one can execute `deactivate` to deactivate the session, or `ctrl + d` combo to quit the terminal.

### 4. Removing the venv
**In a normal usecase one can simply delete the env directory** from the UI after use (yeah for the CLI people, you got `rm -r` option as well).

## Installation

### Install using the `install.sh` script

There is a ready to go shell script one can use to simply install devc and its dependencies into your system. You can find the script over [here](./scripts/install.sh)

#### If you blindly trust the author (@bRuttaZz)
Execute the following to install devc on a single command
```sh
wget https://raw.githubusercontent.com/bRuttaZz/devc/main/scripts/install.sh -O - | sh
```
If you are not a fan of wget, try out curl :)
```sh
curl -o - https://raw.githubusercontent.com/bRuttaZz/devc/main/scripts/install.sh | sh
```

#### Got it you are not the one who blindly trust me
1. Simply download the installation script from [here](https://raw.githubusercontent.com/bRuttaZz/devc/main/scripts/install.sh). (The source can found over [here](https://github.com/bRuttaZz/devc/blob/main/scripts/install.sh))
2. (Optional) Read the script 
3. Make it executable (`chmod +x install.sh`)
4. Run the script in `dry-run` mode (it will not make any changes rather show you how the script is going to behave on your system)
```sh
sh install.sh --dry-run
```
5. If everything seems okay execute 
```sh
sh install.sh 
```
You could have trust me from the first place! JK it's good to analyse a script before its execution

### Building from source

It's too simple,
1. Install the dependencies 
    * buildah
2. Install the build tools
    * go
3. Clone the repo and `cd` into it
4. Execute 
    1. `make` (will create the binary)
    2. (optionally for installing) `make install`
5. Execute `devc --version` to check if everything is doing good.

## Uninstalling
You can simply executing the following to remove the devc binary from your path
```sh
sudo rm $(which devc)  
```
Incase your system have bash installed, devc installation script will add a bash completion tool as well. To remove that execute 
```sh
sudo rm /etc/bash_completion.d/devc-complete
```
In addition devc may have installed `buildah` and it's dependencies. You can simply remove it using your distribution's package manager (apt uninstall, dnf remove, etc..)


## Credits
* Logo : [Dev G](https://www.instagram.com/dev.g.__)
* [Proot](https://github.com/proot-me/proot)
* [Buildah](https://github.com/containers/buildah) 

## Known Issues
* Currently there is some problem regarding issuing apt update in ubuntu/deb based images (gpg key update related issues)
* Some mount points has to be refined. Think some of them are useless at times. (It's hard to generalize between different distros)

## TODOs
* Integration of buildah as a go lib (optional :) 
* packaging (deb, rpm, AppImage) and bash completion
<!-- * logger addition -->

**Plan to be a contributor? You are welcome :)**