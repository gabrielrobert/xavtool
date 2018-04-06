package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/DHowett/go-plist"
	multierror "github.com/hashicorp/go-multierror"
)

type iOSHandler struct {
}

func (h iOSHandler) isPackage(filename string) bool {
	return strings.ToLower(getFilename(filename)) == "info.plist"
}

func (h iOSHandler) getPackageInfo(filePath string) (packageInfo, error) {
	byteValue := readFile(filePath)
	data, err := h.read(byteValue)

	if err != nil {
		return packageInfo{Path: filePath, HasError: true}, err
	}

	return packageInfo{
		Name:            data["CFBundleDisplayName"].(string),
		Version:         data["CFBundleShortVersionString"].(string),
		InternalVersion: data["CFBundleVersion"].(string),
		Path:            filePath,
	}, nil
}

func (h iOSHandler) read(data []byte) (map[string]interface{}, error) {
	var result error

	buffer := bytes.NewReader(data)
	decoder := plist.NewDecoder(buffer)
	var decodeInterface = map[string]interface{}{}
	err := decoder.Decode(&decodeInterface)

	if err != nil {
		return decodeInterface, err
	}

	if _, exists := decodeInterface["CFBundleDisplayName"]; !exists {
		result = multierror.Append(result, fmt.Errorf("Missing property %v", "CFBundleDisplayName"))
	}

	if _, exists := decodeInterface["CFBundleVersion"]; !exists {
		result = multierror.Append(result, fmt.Errorf("Missing property %v", "CFBundleVersion"))
	}

	if _, exists := decodeInterface["CFBundleShortVersionString"]; !exists {
		result = multierror.Append(result, fmt.Errorf("Missing property %v", "CFBundleShortVersionString"))
	}

	return decodeInterface, result
}

func (h iOSHandler) changePackageVersion(file packageInfo, newVersion string) error {
	// open file with all data
	byteValue := readFile(file.Path)
	processedBytes, err := h.applyVersion(byteValue, newVersion)
	if err != nil {
		return fmt.Errorf("Invalid plist file: %v", file.Path)
	}
	saveFile(file.Path, processedBytes)
	return nil
}

func (h iOSHandler) applyVersion(byteValue []byte, newVersion string) ([]byte, error) {
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
