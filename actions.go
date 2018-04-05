package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// called by executing `xavtool current`
func current(c *cli.Context) error {
	allFiles, err := findManifests(getWorkingDir())

	// validations
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if len(allFiles) == 0 {
		return cli.NewExitError("No application has been found", 2)
	}

	for _, file := range allFiles {
		if file.HasError {
			return cli.NewExitError(fmt.Sprintf("Invalid file content:  %v", file.Path), 2)
		}
	}

	// show packages
	for _, file := range allFiles {
		fmt.Println(fmt.Sprintf("%v - %v (%v)", file.Version, file.Name, file.Path))
	}

	return nil
}

// called by executing `xavtool increment`
func increment(c *cli.Context) error {
	allFiles, err := findManifests(getWorkingDir())

	// validations
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if len(allFiles) == 0 {
		return cli.NewExitError("No application has been found", 2)
	}

	for _, file := range allFiles {
		if file.HasError {
			return cli.NewExitError(fmt.Sprintf("Invalid file content:  %v", file.Path), 3)
		}
	}

	// execute version update
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
			return cli.NewExitError(fmt.Sprintf("Invalid type %v", incrementType), 4)
		}

		setVersion(file, newVersion)
		updatedManifest, err := findManifests(file.Path)

		if err != nil {
			return cli.NewExitError(err, 5)
		}

		fmt.Println(fmt.Sprintf("%v: New version: %v (%v)", file.Version, updatedManifest[0].Version, file.Path))
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

	allFiles, err := findManifests(getWorkingDir())

	// validations
	if err != nil {
		return cli.NewExitError(err, 4)
	}

	if len(allFiles) == 0 {
		return cli.NewExitError("No application has been found", 5)
	}

	for _, file := range allFiles {
		if file.HasError {
			return cli.NewExitError(fmt.Sprintf("Invalid file content:  %v", file.Path), 6)
		}
	}

	// execute version update
	for _, file := range allFiles {
		setVersion(file, newVersion)
		updatedManifest, err := findManifests(file.Path)

		if err != nil {
			return cli.NewExitError(err, 7)
		}

		fmt.Println(fmt.Sprintf("%v: New version: %v (%v)", file.Version, updatedManifest[0].Version, file.Path))
	}
	return nil
}

func setVersion(file packageInfo, version string) {
	if isiOsPackage(file.Path) {
		changeiOSPackageVersion(file, version)
	} else if isAndroidPackage(file.Path) {
		changeAndroidPackageVersion(file, version)
	} else if isUWPPackage(file.Path) {
		changeUWPPackageVersion(file, version)
	}
}
