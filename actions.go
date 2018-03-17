package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// called by executing `xavtool current`
func current(c *cli.Context) error {
	allFiles := findManifests(getWorkingDir())
	if len(allFiles) == 0 {
		return cli.NewExitError("No application has been found", 1)
	}

	for _, file := range allFiles {
		fmt.Println(fmt.Sprintf("%v - %v (%v)", file.Version, file.Name, file.Path))
	}

	return nil
}

// called by executing `xavtool increment`
func increment(c *cli.Context) error {
	allFiles := findManifests(getWorkingDir())
	for _, file := range allFiles {

		newVersion := file.Version
		switch incrementType := c.String("type"); incrementType {
		case "major":
			newVersion = incrementMajor(file.Version)
		case "minor":
			newVersion = incrementMinor(file.Version)
		case "patch":
			newVersion = incrementPatch(file.Version)
		default:
			return cli.NewExitError(fmt.Sprintf("Invalid type %v", incrementType), 1)
		}

		setVersion(file, newVersion)
		fmt.Println(fmt.Sprintf("%v: New version: %v (%v)", file.Version, newVersion, file.Path))
	}
	return nil
}

// called by executing `xavtool set`
func set(c *cli.Context) error {
	if !c.Args().Present() {
		return cli.NewExitError("Missing paramter: `version`", 1)
	}
	newVersion := c.Args().Get(0)
	if len(newVersion) <= 0 {
		return cli.NewExitError("Empty parameter: `version`", 2)
	}
	isValid := isVersion(newVersion)
	if !isValid {
		return cli.NewExitError(fmt.Sprintf("Version '%v' is not valid", newVersion), 3)
	}

	allFiles := findManifests(getWorkingDir())
	for _, file := range allFiles {
		setVersion(file, newVersion)
		fmt.Println(fmt.Sprintf("%v: New version: %v (%v)", file.Version, newVersion, file.Path))
	}
	return nil
}

func setVersion(file packageInfo, version string) {
	if isiOsPackage(file.Path) {
		changeiOSPackageVersion(file, version)
	} else if isAndroidPackage(file.Path) {
		changeAndroidPackageVersion(file, version)
	}
}
