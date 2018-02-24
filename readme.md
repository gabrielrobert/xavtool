# Xamarin Automating Version Tool

[![Build status](https://ci.appveyor.com/api/projects/status/6lfimg1j4pw9f807?svg=true)](https://ci.appveyor.com/project/grobert092/xavtool)

Command-line utility to automatically increase iOS / Android / UWP applications version written in [Go](https://golang.org/).

## Installation

From source:

```bash
$ go install
$ xavtool --version
xavtool version *.*.*
```

## Usage

```bash
# Sementic versioning
$ xavtool

NAME:
   xavtool - Command-line utility to automatically increase applications version

USAGE:
   xavtool.exe [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   Gabriel Robert <g.robert092@gmail.com>

COMMANDS:
     current, c  List current versions
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Support

Please [open an issue](https://github.com/gabrielrobert/xavtool/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/gabrielrobert/xavtool/compare).