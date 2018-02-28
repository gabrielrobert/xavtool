package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/clbanning/mxj"
)

type androidBundlerHeader struct {
	XMLName xml.Name `xml:"manifest"`
	Name    string   `xml:"package,attr"`

	// should be less than 2100000000
	Code string `xml:"versionCode,attr"`

	VersionName string     `xml:"versionName,attr"`
	Attrs       []xml.Attr `xml:",attr"`
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

func changeAndroidPackageVersion(file packageInfo, newVersion string) error {
	// open file with all data
	// var data = map[string]interface{}{}
	mv, err := mxj.NewMapXml(readFile(file.Path))

	check(err)
	fmt.Println(mv)
	// increment version
	// data.VersionName = newVersion
	// data.Code = newVersion
	bleh, err := mv.Attributes("android:versionCode")
	fmt.Println(bleh)

	// write data inside buffer
	xmlValue, err := mv.Xml()
	saveFile(file.Path, xmlValue)
	return nil
}
