package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/DHowett/go-plist"
	"github.com/urfave/cli"
)

type Application struct {
	Name    string
	Version string
	OS      string
}

type iOSBundlerHeader struct {
	BundleName    string `plist:"CFBundleDisplayName"`
	BundleVersion string `plist:"CFBundleVersion"`
}

func current(c *cli.Context) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	allFiles := findiOSFile(dir)
	for _, file := range allFiles {
		application := getCurrentVersion(file)
		fmt.Printf("%v [iOS] - %v", application.Name, application.Version)
	}
	return nil
}

func getCurrentVersion(filePath string) Application {
	decoder := plist.NewDecoder(openFile(filePath))
	var data iOSBundlerHeader
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	return Application{Name: data.BundleName, Version: data.BundleVersion}
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
