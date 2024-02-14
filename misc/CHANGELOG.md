# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased] - yyyy-mm-dd
**..**

### Added
### Changed
### Fixed


## [v1.1.1] - 2024-02-14
**Compatibility improvements! Resolution in issues with buildah in ubuntu**

### Added
- Independent build cache handling for devc
### Changed
- Removed the usage of `prune` from buildah. The latest `prune` functionality is implemented using `rmi` flags.
### Fixed
- Support for prune command in devc (now works with older versions)



## [v1.1.0] - 2024-02-14
**Build and Pull based on more stable methods provided by buildah. Removed unnecessary caching on the front end and more.. Now support older versions of buildah as well.**

### Added
- Support for ubuntu
- Now support all foreign pulled images as well as successful image builds persists in the cache
### Changed
- Documentation updates
- Removed `rm` option
- Made the `pull` option more reliable, rather than using an hacked version of build
### Fixed
- Added support for older stable releases of buildah (maintained versions in ubuntu)



## [v1.0.0] - 2024-02-11
**The initial version of *devc*, with its primary features.**

The good news is, if you're on Linux, you can give **devc** a try. It's a streamlined container tool designed for developers working on local development. It's compatible with contemporary container standards like Docker and OCI. **devc** facilitates you with a python-virtual-environment kind of way to interact with your container environments.
