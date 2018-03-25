![Header](_assets/xavtool_header.png "Header")

# Xamarin Automating Version Tool

[![Build status](https://ci.appveyor.com/api/projects/status/6lfimg1j4pw9f807?svg=true)](https://ci.appveyor.com/project/grobert092/xavtool)
[![Build Status](https://travis-ci.org/gabrielrobert/xavtool.svg?branch=master)](https://travis-ci.org/gabrielrobert/xavtool)

Command-line utility to automatically increase iOS / Android / UWP applications version written in [Go](https://golang.org/). It follows [Semantic Versioning](https://semver.org/).

## Installation

### Windows:

Using [Chocolatey](https://chocolatey.org/):

```bash
$ choco install xavtool -version 1.1.0
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
   1.1.0

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

## Typical flow

```bash
$ xavtool current
1.0.1 - androidApp (...\test\AndroidManifest.xml)
1.0.1 - iOSApp (...\test\Info.plist)
1.0.1.0 - uwpApp (...\test\Package.appxmanifest)

$ git flow release start '1.1.0'

$ xavtool i
1.0.1: New version: 1.1.0 (...\test\AndroidManifest.xml)
1.0.1: New version: 1.1.0 (...\test\Info.plist)
1.0.1.0: New version: 1.1.0.0 (...\test\Package.appxmanifest)

$ git commit -am "Version bump to 1.1.0"
$ git flow release finish -p
```

It will update these files:

- `Info.plist`
- `AndroidManifest.xml`
- `Package.appxmanifest`

## Results

### Info.plist

Only these values will be edited:

1) `CFBundleShortVersionString` (new version)
2) `CFBundleVersion` (new version)

Before:
```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <!-- ... -->
        <key>CFBundleShortVersionString</key>
        <string>1.0.1</string>
        <key>CFBundleVersion</key>
        <string>1.0.1</string>
        <!-- ... -->
    </dict>
</plist>
```

After:
```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <!-- ... -->
        <key>CFBundleShortVersionString</key>
        <string>1.1.0</string>
        <key>CFBundleVersion</key>
        <string>1.1.0</string>
        <!-- ... -->
    </dict>
</plist>
```

### AndroidManifest.xml

Only these values will be edited:

1) `manifest/@android:versionName` (new version)
2) `manifest/@android:versionCode` (integer computed this way: `(major * 1000000) + (minor * 10000) + (patch * 100)`)        

Before:
```xml
<?xml version="1.0" encoding="utf-8"?>
<manifest
    xmlns:android="http://schemas.android.com/apk/res/android" 
    package="com.example.xavtool" 
    android:versionCode="1000100"
    android:versionName="1.0.1">
    <!-- ... -->
</manifest>
```

After:
```xml
<?xml version="1.0" encoding="utf-8"?>
<manifest
    xmlns:android="http://schemas.android.com/apk/res/android" 
        package="com.example.xavtool" 
        android:versionCode="1010000" 
        android:versionName="1.1.0">
    <!-- ... -->
</manifest>
```

### Package.appxmanifest

Only these values will be edited:

1) `Package/Identity/@Version` (new version)

Before:
```xml
<?xml version="1.0" encoding="utf-8"?>
<Package
    xmlns="http://schemas.microsoft.com/appx/manifest/foundation/windows10"
    xmlns:mp="http://schemas.microsoft.com/appx/2014/phone/manifest"
    xmlns:uap="http://schemas.microsoft.com/appx/manifest/uap/windows10" IgnorableNamespaces="uap mp">
    
    <!-- ... -->
    <Identity Name="95748d56-342b-4dae-93f5-aeda0587a1c0" Publisher="CN=gabrielrobert" Version="1.0.1"/>
    <!-- ... -->
    
</Package>
```

After:
```xml
<?xml version="1.0" encoding="utf-8"?>
<Package
    xmlns="http://schemas.microsoft.com/appx/manifest/foundation/windows10"
    xmlns:mp="http://schemas.microsoft.com/appx/2014/phone/manifest"
    xmlns:uap="http://schemas.microsoft.com/appx/manifest/uap/windows10" IgnorableNamespaces="uap mp">
    
    <!-- ... -->
    <Identity Name="95748d56-342b-4dae-93f5-aeda0587a1c0" Publisher="CN=gabrielrobert" Version="1.1.0"/>
    <!-- ... -->
    
</Package>
```

## Support

Please [open an issue](https://github.com/gabrielrobert/xavtool/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/gabrielrobert/xavtool/compare).