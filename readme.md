![Header](_assets/xavtool_header.png "Header")

# Xamarin Automating Version Tool

[![Build status](https://ci.appveyor.com/api/projects/status/6lfimg1j4pw9f807?svg=true)](https://ci.appveyor.com/project/grobert092/xavtool)
[![Build Status](https://travis-ci.org/gabrielrobert/xavtool.svg?branch=master)](https://travis-ci.org/gabrielrobert/xavtool)

Command-line utility to automatically increase iOS / Android / UWP applications version written in [Go](https://golang.org/). It follows [Semantic Versioning](https://semver.org/).

## Installation

### Windows:

Using [Chocolatey](https://chocolatey.org/):

```bash
$ choco install xavtool
$ xavtool --version
```

Using [scoop](http://scoop.sh/):

```bash
$ scoop bucket add gabrielrobert-bucket https://github.com/gabrielrobert/scoop-bucket
$ scoop install xavtool
```

### macOS:

Using [brew](https://brew.sh/):

```bash
$ brew install gabrielrobert/tap/xavtool
$ xavtool --version
```

### Binaries

Download executables on the [release page](https://github.com/gabrielrobert/xavtool/releases/latest).

### From source:

```bash
$ go build
$ go test -v
$ go install
$ xavtool --version
```

## Usage

```bash
$ xavtool

NAME:
   xavtool - Command-line utility to automatically increase applications version

USAGE:
   xavtool [global options] command [command options] [arguments...]

VERSION:
   0.11.2

AUTHOR:
   Gabriel Robert <g.robert092@gmail.com>

COMMANDS:
     current, c    List current versions
     increment, i  Increment to next version
     set, s        Set the current project version
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Typical flow

```bash
$ git flow release start '1.16.0'
$ xavtool i
$ git commit -am "Version bump to 1.16.0"
$ git flow release finish -p
```

### increment
```bash
$ xavtool increment --help

NAME:
   xavtool increment - Increment to next version

USAGE:
   xavtool increment [command options] [arguments...]

OPTIONS:
   --type value, -t value  major, minor, path (default: "minor")
```

### set
```bash
$ xavtool set --help

NAME:
   xavtool set - Set the current project version

USAGE:
   xavtool set [arguments...]
```

## Support

Please [open an issue](https://github.com/gabrielrobert/xavtool/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/gabrielrobert/xavtool/compare).