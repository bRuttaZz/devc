# devC

<a href="./assets/LICENSE.md" title="Logo by Dev G"><img src="./assets/tad.png" width=150 alt="Logo by Dev G"/></a>

Containers For Developers (Container as a Folder)

![Workflow status](https://github.com/bruttazz/devc/actions/workflows/codeql.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/bruttazz/devc)](https://goreportcard.com/report/github.com/bruttazz/devc)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://github.com/bruttazz/devc/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/bruttazz/devc.svg)](https://pkg.go.dev/github.com/bruttazz/devc)
[![Release](https://img.shields.io/github/release/bruttazz/devc.svg?style=flat-square)](https://github.com/bruttazz/devc/releases/latest)

---

## Intro

> **TL;DR**: The good news is, if you're on Linux, you can give **devc** a try. It's a streamlined container tool designed for developers working on local development, compatible with contemporary container standards like Docker and OCI. **devc** presents you with a python-virtualenvironment kind of way to interact with container environment.  


Have you ever utilized Docker or Podman in local application development, just for resolving the dependency issues? If so, have you encountered challenges in creating and maintaining the Docker environment for development? This may involve rebuilding your project with each Dockerfile change, mounting your local codebase to the container, dealing with permission overhead, and connecting your IDE to the container.

The good news is, if you're on Linux, you can give **devc** a try. It's a streamlined container tool designed for developers working on local development, compatible with contemporary container standards like Docker and OCI. **devc** presents you with a python-virtualenvironment kind of way to interact with container environment. 
Basically, using devc one can create a **new venv similar to python-venv** from existing Dokckerfile/Containerfile or by specifying a docker or oci image at a registry (eg: docker-hub). After creation the developer can activate that env using `source venv/bin/activate` (yes, exactly like python venv). And (if everything went well) from the next instance your termnial prompt will be prefixed with `(devc)` and you will be inside the newly created container environment, with all your files and directories from the current working directory. The venv may of different linux distribution with different setup, whatever be that, one can install any application they want to it, without affecting the host machine. After you are tired up by playing with it, simply type `deactivate` (yeah you got it), or `ctrl + d` keyboard combination to exit from the env, and you will be on your normal original terminal. Want to enter again? type the same command `source venv/bin/activate`, all your previously made changes will be waiting over there. And at the end of the day. You can simply delete the venv directory. and that's how you going to manage the storage (just like in the good old days before the storage hungry dockers).

**NB** : In case you missed it, **devc** is just a simple developer tool, it's meant to be used for local development, and it's a striped down version of normal container setups. 

## Yet Another Container Tool ?
Okay you may be thinking **devc** as a wrap around bunch of Docker or Podman workflows to present it similar to the way of python venv. But infact devc handles things in a different way. Docker, Podman or other similar container tools are optimised to work with deployment time, and they got a lot of process management and resource sharing features as well. But in devc we are taking only the sandboxed rootfs and binding system. Following points may detail about the behavior of devc
1. Devc will be maintaining a standalone *rootfs*, so that the users will be able to install and run different programs without affecting the host system, other than were the current working directory is situated.
2. Unlike other container tools, Devc does not have any deamon processes of its own. Every command issued on devc environment will be run as in host namespace. Making it lightweight, and developer friendly.
3. Devc can run completely on user space. It's absolutely **not recommended to run devc with superuser privillages**
4. The container build cache and the intermediate images pulled while building a devc environment will be stored under the venv (although the cache directories will be wiped as soon as the container images are built). Deleting it will delete all other cache dependent upon the venv.
5. The current working directoy on activating the devc environment will be shared with the default workdir of devc (All other file system will be in a sandboxed environment). Due to this one can seamlessly make changes in the files in current work directory and try out the changes in the activated devc container environment.

**About the Name** : As you may have guessed by now, the name "devc" is made of the two words "developer" and "container". i.e., **containers for developers**.

## Usage

## Installation

## TODOs
* Integration of buildah as a go lib (optional :) 
* logger addition
* packaging (deb, rpm)
* Complete the docs, create script and make ci cd

<br><br>

**Credits**
* Logo : [Dev G](https://www.instagram.com/dev.g.__)
* [Proot](https://github.com/proot-me/proot)
* [Buildah](https://github.com/containers/buildah) 