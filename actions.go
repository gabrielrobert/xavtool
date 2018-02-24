package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DHowett/go-plist"
	"github.com/urfave/cli"
)

type application struct {
	Name    string
	Version string
	Path    string
}

type iOSBundlerHeader struct {
	BundleName    string `plist:"CFBundleDisplayName"`
	BundleVersion string `plist:"CFBundleVersion"`
}

func current(c *cli.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	allFiles := findiOSFile(dir)
	for _, file := range allFiles {
		info := getCurrentVersion(file)
		fmt.Println(fmt.Sprintf("%v - [iOS] - %v (%v)", info.Version, info.Name, info.Path))
	}
	return nil
}

func getCurrentVersion(filePath string) application {
	decoder := plist.NewDecoder(openFile(filePath))
	var data iOSBundlerHeader
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	return application{
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
