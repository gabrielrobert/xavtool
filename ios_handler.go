package main

import (
	"fmt"
	"strings"

	"github.com/DHowett/go-plist"
)

type iOSBundlerHeader struct {
	BundleName    string `plist:"CFBundleDisplayName"`
	BundleVersion string `plist:"CFBundleVersion"`
}

func isiOsPackage(filename string) bool {
	return strings.ToLower(filename) == "info.plist"
}

func getiOSPackageInfo(filePath string) packageInfo {
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
