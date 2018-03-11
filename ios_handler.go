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
	byteValue := readFile(filePath)
	data := readiOSData(byteValue)
	return packageInfo{
		Name:    data["CFBundleDisplayName"].(string),
		Version: data["CFBundleVersion"].(string),
		Path:    filePath,
	}
}

func readiOSData(data []byte) map[string]interface{} {
	buffer := bytes.NewReader(data)
	decoder := plist.NewDecoder(buffer)
	var decodeInterface = map[string]interface{}{}
	err := decoder.Decode(&decodeInterface)
	check(err)
	return decodeInterface
}

func changeiOSPackageVersion(file packageInfo, newVersion string) error {
	// open file with all data
	byteValue := readFile(file.Path)
	processedBytes := applyVersionToiOSPlist(byteValue, newVersion)
	saveFile(file.Path, processedBytes)
	return nil
}

func applyVersionToiOSPlist(byteValue []byte, newVersion string) []byte {
	buffer := bytes.NewReader(byteValue)
	decoder := plist.NewDecoder(buffer)
	var data = map[string]interface{}{}
	err := decoder.Decode(&data)
	check(err)

	// increment version
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

	return bufferedData.Bytes()
}
