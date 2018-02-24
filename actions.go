package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DHowett/go-plist"
	"github.com/urfave/cli"
)

type iOSBundlerHeader struct {
	BundleName    string `plist:"CFBundleDisplayName"`
	BundleVersion string `plist:"CFBundleVersion"`
}

func current(c *cli.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	allFiles := findManifests(dir)
	for _, file := range allFiles {
		info := getCurrentVersion(file)
		fmt.Println(fmt.Sprintf("%v - %v (%v)", info.Version, info.Name, info.Path))
	}
	return nil
}

func getCurrentVersion(filePath string) packageInfo {
	decoder := plist.NewDecoder(openFile(filePath))
	var data iOSBundlerHeader
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	return packageInfo{
		Name:    data.BundleName,
		Version: data.BundleVersion,
		Path:    filePath,
	}
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
