package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// called by executing `xavtool current`
func current(c *cli.Context) error {
	dir := getWorkingDir()
	allFiles := findManifests(dir)
	if len(allFiles) == 0 {
		fmt.Println("No application has been found")
	} else {
		for _, file := range allFiles {
			fmt.Println(fmt.Sprintf("%v - %v (%v)", file.Version, file.Name, file.Path))
		}
	}

	return nil
}

// called by executing `xavtool increment`
func increment(c *cli.Context) error {
	dir := getWorkingDir()
	allFiles := findManifests(dir)
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
			fmt.Println(fmt.Sprintf("Invalid type %v", incrementType))
		}

		changeiOSPackageVersion(file, newVersion)
		fmt.Println(fmt.Sprintf("%v: New version: %v (%v)", file.Version, newVersion, file.Path))
	}
	return nil
}
