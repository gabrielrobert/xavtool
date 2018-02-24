package main

import (
	"fmt"
	"os"

	"github.com/DHowett/go-plist"
	"github.com/urfave/cli"
)

type iOSBundlerHeader struct {
	BundleVersion string `plist:"CFBundleVersion"`
}

func current(c *cli.Context) error {
	fmt.Printf("iOS version - %v", getCurrentVersion("test/Info.plist"))
	return nil
}

func getCurrentVersion(filePath string) string {
	file := openFile(filePath)
	decoder := plist.NewDecoder(file)
	var data iOSBundlerHeader
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	return data.BundleVersion
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
