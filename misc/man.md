# DEVC 
[![Release](https://img.shields.io/github/release/bruttazz/devc.svg?style=flat-square)](https://github.com/bruttazz/devc/releases/latest)

**devc: Containers for Developers (Container as a Directory)!**
Offers developers a user-friendly interface for constructing applications using containers! Generate effortlessly usable virtual environments from Containerfiles or container images, activate them, engage in development, implement changes, and reuse seamlessly


**Usage**:
  * `devc [flags]`
  * `devc [command]`

**Flags**:
  * `-h`, `--help` :  help for devc
  * `-v`, `--version` :  Get current version

**Available Commands**:
  * `activate` :   activate a devc env
  * `build`   :    build a devc env from a Dockerfile or Containerfile
  * `completion`:  Generate the autocompletion script for the specified shell
  * `help`       : Help about any command
  * `images`      :show all the cached images
  * `login`       :login into a container-image registry
  * `logout`      :logout from a previously * logged in registry
  * `prune`       :prune all the cached images.
  * `pull`        :build a devc env directly from a container image
  * `rm`          :remove a devc environment
  * `rmi`         :remove a cached image


### activate
**activate a devc environment**. 

One can also enable the env by executing `source <env-name>/bin/activate`.
On activating the env, it will automatically mount the current working
dir to an unused directory (/< a randomly created uuid>/devc) inside the environment

**Usage**:
  * `devc activate env-name [flags]`

**Flags**:
  * `-h`, `--help`:                 help for activate
  * `-w`, `--working-dir` : [string]   Use a deferent working directory on activation, defaults to '/devc'

### build
**build an environment from a Dockerfile or a Containerfile.**

The file is intended to be available from the current working directory. 
Otherwise specify the location of the file with "--file" flag.
If no specific arguments are provided, will use the current working directory as the 
build context and look for a Docker/Containerfile in it. The build fails if no
Containerfile or Dockerfile is present.

**Usage**:
  * `devc build [options] [flags] env-name`

**Flags**:	
* `--build-arg`: [string]  argument=value to supply build time arguments to the builder
* `-c`, `--context`: [string]             Specify the build context to be used (default ".")
* `-f`, `--file`: [string]                Explicitly point to a container file (can also provide valid urls)
* `-h`, `--help`: help for build
* `--keep-cache` : Keep build cache after successful build of devc environment. By default the build-cache will not be stored. **NB** : if a devc env is created using `--keep-cache` flag, try to use `devc rm <env-name>` to remove the env after use.

### images
**Show all the cache images.** 

Images will be cached for latter use if devc env is created using `pull` command. See `prune` and `rmi` commands to remove the cached images.

**Usage**:
  * `devc images [flags]`

**Flags**:
  `-h`, `--help` :  help for images

### login
**login into a container-image registry, using the username and password.**

Can be use to pull images from private registries. "registry-uri" can be the domain name or domain with port number. eg: `registry.hub.docker.com`

**Usage**:
  * `devc login registry-uri [flags]`

**Flags**:
  * `-h`, `--help` : help for login
  * `--password`: [string]   Optionally specify password. if not provided will be propted from via stdin
  * `--username`: [string]   Optionally specify username. If not provided will be propted from via stdin

### logout
**logout from a previously logged in registries using "devc login".**

**Usage**:
  * `devc logout registry-uri [flags]`

**Flags**:
  * `-h`, `--help` :  help for logout


### prune
**Remove all the cached images.** 

Images will be cached for latter use if devc env is created using `pull` command.
Using `--wipe` option will clear everything including the registry login credentials

**Usage**:
  * `devc prune [flags]`

**Flags**:
  * `-h`, `--help`  : help for prune
  * `--wipe`:   Clear all cache from user session, including registry login credentials!

### pull
**build an devc environment from a container image.**

The first argument 'image-name' (optionally with a tag) will be search inside
available registries. One can add new registries by executing `devc login`

**Usage**:
  * `devc pull [options] image-name env-name [flags]`

**Flags**:
  * `-h`, `--help` :  help for pull
    * `--rm` :     Do not create local image cache, after building devc environment

### rm
**remove the specified devc environment.** 

Can be used if it's not able to remove the directory otherwise, or you have issued special flags like `--keep-cache` on building the devc env.

**Usage**:
  * `devc rm env-name [flags]`

**Flags**:
  * `-h`, `--help` :   help for rm

### rmi
**remove a cached image by id**

Images will be cached for latter use if devc env is created using `pull` command. See `devc images` to see all the cached images.

**Usage**:
  * `devc rmi image-id [flags]`

**Flags**:
  * `-h`, `--help` :   help for rmi

