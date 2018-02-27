package main

import (
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

func findManifests(root string) []packageInfo {
	fileList := []packageInfo{}

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {

		if toBeIgnored(f) {
			return filepath.SkipDir
		}

		if isManifestFile(f) {
			if isiOsPackage(f.Name()) {
				fileList = append(fileList, getiOSPackageInfo(path))
			} else if isAndroidPackage(f.Name()) {
				fileList = append(fileList, getAndroidPackageInfo(path))
			}
		}
		return nil
	})

	check(err)
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
