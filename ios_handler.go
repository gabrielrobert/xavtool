package main

import (
	"bytes"
	"encoding/binary"
	"strings"

	"github.com/DHowett/go-plist"
)

func isiOsPackage(filename string) bool {
	return strings.ToLower(filename) == "info.plist"
}

func getiOSPackageInfo(filePath string) packageInfo {
	decoder := plist.NewDecoder(openFile(filePath))
	var data = map[string]interface{}{}
	err := decoder.Decode(&data)
	check(err)
	return packageInfo{
		Name:    data["CFBundleDisplayName"].(string),
		Version: data["CFBundleVersion"].(string),
		Path:    filePath,
	}
}

func changeiOSPackageVersion(file packageInfo, newVersion string) error {
	// open file with all data
	decoder := plist.NewDecoder(openFile(file.Path))
	var data = map[string]interface{}{}
	err := decoder.Decode(&data)
	data["CFBundleVersion"] = newVersion
	data["CFBundleShortVersionString"] = newVersion

	// write data inside buffer
	var bufferedData bytes.Buffer
	binary.Write(&bufferedData, binary.BigEndian, data)

	// encode data
	encoder := plist.NewEncoder(&bufferedData)
	encoder.Indent("\t")
	err = encoder.Encode(data)
	check(err)

	saveFile(file.Path, bufferedData.Bytes())
	return nil
}
