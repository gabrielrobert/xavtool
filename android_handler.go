package main

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
)

type androidBundlerHeader struct {
	XMLName xml.Name `xml:"manifest"`
	Name    string   `xml:"package,attr"`

	// should be less than 2100000000
	Code string `xml:"versionCode,attr"`

	VersionName string `xml:"versionName,attr"`
}

func isAndroidPackage(filename string) bool {
	return strings.ToLower(filename) == "androidmanifest.xml"
}

func getAndroidPackageInfo(filePath string) packageInfo {
	byteValue, _ := ioutil.ReadAll(openFile(filePath))
	var data androidBundlerHeader
	xml.Unmarshal(byteValue, &data)
	return packageInfo{
		Name:    data.Name,
		Version: data.VersionName,
		Path:    filePath,
	}
}
