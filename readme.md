# Xamarin Automating Version Tool

[![Build status](https://ci.appveyor.com/api/projects/status/6lfimg1j4pw9f807?svg=true)](https://ci.appveyor.com/project/grobert092/xavtool)
[![Build Status](https://travis-ci.org/gabrielrobert/xavtool.svg?branch=master)](https://travis-ci.org/gabrielrobert/xavtool)

Command-line utility to automatically increase iOS / Android / UWP applications version written in [Go](https://golang.org/).

## Installation

### Windows:

Using chocolatey:

```bash
$ choco install xavtool
$ xavtool --version
```

### macOS:

Using brew:

```bash
$ brew install gabrielrobert/tap/xavtool
$ xavtool --version
```

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
   0.1.0

AUTHOR:
   Gabriel Robert <g.robert092@gmail.com>

COMMANDS:
     current, c    List current versions
     increment, i  Increment to next version
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
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

## Support

Please [open an issue](https://github.com/gabrielrobert/xavtool/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/gabrielrobert/xavtool/compare).