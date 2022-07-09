<div align="center">
	<img src="./assets/hilbish-flower.png" width=128><br>
	<img src="./assets/hilbish-text.png" width=256><br>
	<blockquote>
	🌺 The flower shell. A comfy and nice little shell for Lua fans!
	</blockquote>
	<p align="center">
		<img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/Rosettea/Hilbish?style=flat-square">
		<img alt="GitHub commits since latest release (by date)" src="https://img.shields.io/github/commits-since/Rosettea/Hilbish/latest?style=flat-square">
		<img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Rosettea/Hilbish?style=flat-square"><br>
		<a href="https://github.com/Rosettea/Hilbish/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22"><img src="https://img.shields.io/github/issues/Hilbis/Hilbish/help%20wanted?style=flat-square&color=green" alt="help wanted"></a>
		<a href="https://github.com/Rosettea/Hilbish/blob/master/LICENSE"><img alt="GitHub license" src="https://img.shields.io/github/license/Rosettea/Hilbish?style=flat-square"></a>
		<a href="https://discord.gg/3PDdcQz"><img alt="Discord" src="https://img.shields.io/discord/732357621503229962?color=blue&style=flat-square"></a>
	</p>
</div>

Hilbish is a extensible shell (framework). It was made to be very customizable
via the Lua programming language. It aims to be easy to use for the casual
people but powerful for those who want to tinker more with their shell,
the thing used to interface with most of the system.  

The motivation for choosing Lua was that its simpler and better to use
than old shell script. It's fine for basic interactive shell uses,
but that's the only place Hilbish has shell script; everything else is Lua
and aims to be infinitely configurable. If something isn't, open an issue!

# Table of Contents
- [Screenshots](#Screenshots)
- [Installation](#Installation)
  - [Prebuilt Bins](#Prebuilt-binaries)
  - [AUR](#AUR)
  - [Nixpkgs](#Nixpkgs)
  - [Manual Build](#Manual-Build)
- [Getting Started](#Getting-Started)
- [Contributing](#Contributing)

# Screenshots
<div align="center">
<img src="gallery/default.png"><br><br>
<img src="gallery/terminal.png"><br><br>
<img src="gallery/pillprompt.png">
</div>

# Installation
## Prebuilt binaries
Go [here](https://nightly.link/Rosettea/Hilbish/workflows/build/master) for
builds on the master branch.

## AUR
[![AUR maintainer](https://img.shields.io/aur/maintainer/hilbish?logo=arch-linux&style=flat-square)](https://aur.archlinux.org/packages/hilbish)  
Arch Linux users can install Hilbish from the AUR with the following command:  
```sh
yay -S hilbish
```

[![AUR maintainer](https://img.shields.io/aur/maintainer/hilbish?logo=arch-linux&style=flat-square)](https://aur.archlinux.org/packages/hilbish-git)  
Or from the latest `master` commit with:  
```sh
yay -S hilbish-git
```

## Nixpkgs
Nix/NixOS users can install Hilbish from the central repository, nixpkgs, through the usual ways.
If you're new to nix you should probably read up on how to do that [here](https://nixos.wiki/wiki/Cheatsheet).

## Manual Build
### Prerequisites
- [Go 1.17+](https://go.dev)
- [Task](https://taskfile.dev/#/)

### Build
First, clone Hilbish. The recursive is required, as some Lua libraries
are submodules.  
```sh
git clone --recursive https://github.com/Rosettea/Hilbish
cd Hilbish
go get -d ./...
```  

To build, run:
```
task
```  

Or, if you want a stable branch, run these commands:
```
git checkout $(git describe --tags `git rev-list --tags --max-count=1`)
task build
```  

After you did all that, run `sudo task install` to install Hilbish globally.

# Contributing
Any kind of contributions are welcome! Hilbish is very easy to contribute to.
Read [CONTRIBUTING.md](CONTRIBUTING.md) as a guideline to doing so.

**Thanks to everyone below who's contributed!**  
<a href="https://github.com/Rosettea/Hilbish/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Rosettea/Hilbish" />
</a>

*Made with [contributors-img](https://contrib.rocks).*

# License
Hilbish is licensed under the [MIT license](LICENSE).  
[Images and assets](assets/) are licensed under CC-BY-SA 4.0
