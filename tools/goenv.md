Goup
goup (pronounced Go Up) is an elegant Go version manager.

It is notorious that an operating system's package manager takes time to update Go to the latest version, particularly on Linux distros. At the time of this writing in October 2020, Fedora 32's Go version from dnf is 1.14.9, while the latest Go version is 1.15.2.

There are a bunch of solutions to install Go or manage Go versions outside of a package manager: golang/dl, getgo, gvm, goenv, to name a few. All of them either do not work well on all Linux distros (I ran into errors with gvm and goenv on Fedora) or do not provide the developer experience that I like (golang/dl requires a Go compiler to pre-exist; getgo can only install the latest Go)

I want a Go version manager that:

Has a minimum prerequisite to install, e.g., does not need a Go compiler to pre-exist.
Is installed with a one-liner.
Runs well on all operating systems (at least runs well on *uix as a start).
Installs any version of Go (any version from golang.org/dl or tip) and switches to it.
Does not inject magic into your shell.
Is written in Go.
goup is an attempt to fulfill the above features and is heavily inspired by Rustup, golang/dl and getgo.

Installation
One-liner
curl -sSf https://raw.githubusercontent.com/owenthereal/goup/master/install.sh | sh
Install by skipping the confirmation prompt, e.g., for automation:

curl -sSf https://raw.githubusercontent.com/owenthereal/goup/master/install.sh | sh -s -- '--skip-prompt'
Manual
If you want to install manually, there are the steps:

Download the latest goup from https://github.com/owenthereal/goup/releases
Drop the goup executable to your PATH and make it executable: mv GOUP_BIN /usr/local/bin/goup && chmod +x /usr/local/bin/goup
Add the Go bin directory to your shell startup script: echo 'export PATH="$HOME/.go/current/bin:$PATH"' >> ~/.bashrc
Quick Start
$ goup install
Downloaded   0.0% (    32768 / 121149509 bytes) ...
Downloaded  12.4% ( 15007632 / 121149509 bytes) ...
Downloaded  30.2% ( 36634352 / 121149509 bytes) ...
Downloaded  47.6% ( 57703440 / 121149509 bytes) ...
Downloaded  65.9% ( 79855008 / 121149509 bytes) ...
Downloaded  84.2% (101972672 / 121149509 bytes) ...
Downloaded 100.0% (121149509 / 121149509 bytes)
INFO[0030] Unpacking /home/owen/.go/go1.15.2/go1.15.2.linux-amd64.tar.gz ...
INFO[0043] Success: go1.15.2 downloaded in /home/owen/.go/go1.15.2
INFO[0043] Default Go is set to 'go1.15.2'
$ goup ls
go1.15.2
$ go env GOROOT
/home/owen/.go/go1.15.2
$ go version
go version go1.15.2 linux/amd64

$ goup install tip
Cloning into '/home/owen/.go/gotip'...
remote: Counting objects: 10041, done
remote: Finding sources: 100% (10041/10041)
remote: Total 10041 (delta 1347), reused 6538 (delta 1347)
Receiving objects: 100% (10041/10041), 23.83 MiB | 3.16 MiB/s, done.
Resolving deltas: 100% (1347/1347), done.
Updating files: 100% (9212/9212), done.
INFO[0078] Updating the go development tree...
From https://go.googlesource.com/go
 * branch            master     -> FETCH_HEAD
HEAD is now at 5d13781 cmd/cgo: add more architectures to size maps
Building Go cmd/dist using /home/owen/.go/go1.15.2. (go1.15.2 linux/amd64)
Building Go toolchain1 using /home/owen/.go/go1.15.2.
Building Go bootstrap cmd/go (go_bootstrap) using Go toolchain1.
Building Go toolchain2 using go_bootstrap and Go toolchain1.
Building Go toolchain3 using go_bootstrap and Go toolchain2.
Building packages and commands for linux/amd64.
---
Installed Go for linux/amd64 in /home/owen/.go/gotip
Installed commands in /home/owen/.go/gotip/bin
INFO[0297] Default Go is set to 'gotip'
$ goup ls
gotip
$ go env GOROOT
/home/owen/.go/gotip
$ go version
go version devel +5d13781 Thu Oct 8 00:28:09 2020 +0000 linux/amd64


$ GOUP_GO_HOST=golang.google.cn goup install # For Gophers in China, see https://github.com/owenthereal/goup/issues/2
How it works
install.sh downloads the latest Goup release for your platform and appends Goup's bin directory ($HOME/.go/bin) & Go's bin directory ($HOME/.go/current/bin) to your PATH environment variable.
goup switches to selected Go version.
goup set switches to selected Go version.
goup install downloads specified version of Go to$HOME/.go/VERSION and symlinks it to $HOME/.go/current.
goup ls list all installed Go version located at $HOME/.go/current.
goup remove removes the specified Go version.
goup search lists all available Go versions from https://golang.org/dl.
goup upgrade upgrades goup.


资源地址是在： https://github.com/owenthereal/goup

安装之后会输出
Welcome to Goup!

Goup and Go will be located at:

  /root/.go

The Goup command will be located at:

  /root/.go/bin

The go, gofmt and other Go commands will be located at:

  /root/.go/current/bin

To get started you need Goup's bin directory (/root/.go/bin) and
Go's bin directory (/root/.go/current/bin) in your PATH environment
variable. These two paths will be added to your PATH environment variable by
modifying the profile files located at:

  /root/.profile
  /root/.zprofile
  /root/.bash_profile

Next time you log in this will be done automatically. To configure your
current shell run source $HOME/.go/env.