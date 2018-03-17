package main

import (
	"os"
	"path/filepath"
)

type packageInfo struct {
	Name     string
	Version  string
	Path     string
	HasError bool
}

func findManifests(root string) []packageInfo {
	fileList := []packageInfo{}

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {

		if toBeIgnored(f) {
			return filepath.SkipDir
		}

		if isiOsPackage(f.Name()) {
			fileList = append(fileList, getiOSPackageInfo(path))
		} else if isAndroidPackage(f.Name()) {
			fileList = append(fileList, getAndroidPackageInfo(path))
		} else if isUWPPackage(f.Name()) {
			fileList = append(fileList, getUWPPackageInfo(path))
		}

		return nil
	})

	check(err)
	return fileList
}

func toBeIgnored(f os.FileInfo) bool {
	if f.IsDir() && stringInSlice(f.Name(), []string{"bin", "obj", ".git"}) {
		return true
	}
	return false
}
