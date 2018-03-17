package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/DHowett/go-plist"
)

func isiOsPackage(filename string) bool {
	return strings.ToLower(getFilename(filename)) == "info.plist"
}

func getiOSPackageInfo(filePath string) packageInfo {
	byteValue := readFile(filePath)
	data, err := readiOSData(byteValue)

	if err != nil {
		return packageInfo{Path: filePath, HasError: true}
	}

	return packageInfo{
		Name:    data["CFBundleDisplayName"].(string),
		Version: data["CFBundleVersion"].(string),
		Path:    filePath,
	}
}

func readiOSData(data []byte) (map[string]interface{}, error) {
	buffer := bytes.NewReader(data)
	decoder := plist.NewDecoder(buffer)
	var decodeInterface = map[string]interface{}{}
	err := decoder.Decode(&decodeInterface)
	return decodeInterface, err
}

func changeiOSPackageVersion(file packageInfo, newVersion string) error {
	// open file with all data
	byteValue := readFile(file.Path)
	processedBytes, err := applyVersionToiOSPlist(byteValue, newVersion)
	if err != nil {
		return fmt.Errorf("Invalid plist file: %v", file.Path)
	}
	saveFile(file.Path, processedBytes)
	return nil
}

func applyVersionToiOSPlist(byteValue []byte, newVersion string) ([]byte, error) {
	buffer := bytes.NewReader(byteValue)
	decoder := plist.NewDecoder(buffer)
	var data = map[string]interface{}{}
	err := decoder.Decode(&data)

	if err != nil {
		return nil, errors.New("Invalid plist")
	}

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

	return bufferedData.Bytes(), nil
}
