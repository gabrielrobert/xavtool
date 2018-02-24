package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type packageInfo struct {
	Name    string
	Version string
	Path    string
}

var files = []string{"info.plist", "androidmanifest.xml"}
var ignoreFolders = []string{"bin", "obj", ".git"}

func findManifests(root string) []string {
	fileList := []string{}

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {

		if toBeIgnored(f) {
			return filepath.SkipDir
		}

		if isManifestFile(f) {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return fileList
}

func toBeIgnored(f os.FileInfo) bool {
	if f.IsDir() && stringInSlice(f.Name(), ignoreFolders) {
		return true
	}
	return false
}

func isManifestFile(f os.FileInfo) bool {
	loweredFileName := strings.ToLower(f.Name())
	if !f.IsDir() && stringInSlice(loweredFileName, files) {
		return true
	}
	return false
}
