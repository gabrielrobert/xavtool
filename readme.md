# Xamarin Automating Version Tool

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
$ xavtool increment major
$ xavtool increment minor
$ xavtool increment path

$ xavtool downgrade major
$ xavtool downgrade minor
$ xavtool downgrade path

# Infos
xavtool current
```

## Support

Please [open an issue](https://github.com/gabrielrobert/xavtool/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/gabrielrobert/xavtool/compare).