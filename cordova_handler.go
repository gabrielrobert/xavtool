package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/clbanning/mxj"
)

type cordovaHandler struct {
}

type cordovaBundlerHeader struct {
	XMLName xml.Name `xml:"widget"`
	Name    string   `xml:"id,attr"`

	// should be less than 2100000000
	Code string `xml:"android-versionCode,attr"`

	VersionName string     `xml:"version,attr"`
	Attrs       []xml.Attr `xml:",attr"`
}

func (h cordovaHandler) isPackage(filename string) bool {
	return strings.ToLower(getFilename(filename)) == "config.xml"
}

func (h cordovaHandler) getPackageInfo(filePath string) (packageInfo, error) {
	byteValue := readFile(filePath)
	data, err := h.read(byteValue)

	if err != nil {
		return packageInfo{Path: filePath, HasError: true}, err
	}

	return packageInfo{
		Name:            data.Name,
		Version:         data.VersionName,
		InternalVersion: data.Code,
		Path:            filePath,
	}, nil
}

func (h cordovaHandler) read(data []byte) (*cordovaBundlerHeader, error) {
	var header cordovaBundlerHeader
	err := xml.Unmarshal(data, &header)
	return &header, err
}

func (h cordovaHandler) changePackageVersion(file packageInfo, newVersion string) error {
	fileBytes := readFile(file.Path)
	processedBytes, err := h.applyVersion(fileBytes, newVersion)
	if err != nil {
		return fmt.Errorf("Invalid xml file: %v", file.Path)
	}
	saveFile(file.Path, processedBytes)
	return nil
}

func (h cordovaHandler) applyVersion(byteValue []byte, newVersion string) ([]byte, error) {
	fileReader := bytes.NewReader(byteValue)
	for m, err := mxj.NewMapXmlSeqReader(fileReader); m != nil || err != io.EOF; m, err = mxj.NewMapXmlSeqReader(fileReader) {
		if err != nil {
			if err == mxj.NO_ROOT {
				continue
			} else {
				return nil, errors.New("Invalid xml")
			}
		}
		vmap := m["widget"].(map[string]interface{})

		// edit version attr
		acmt, err := mxj.Map(vmap).ValueForPath("#attr.version.#text")
		acmt = newVersion
		mxj.Map(vmap).SetValueForPath(acmt, "#attr.version.#text")
		err = m.SetValueForPath(vmap, "widget")
		check(err)

		// edit ios-CFBundleVersion attr
		acmt, err = mxj.Map(vmap).ValueForPath("#attr.ios-CFBundleVersion.#text")
		acmt = newVersion
		mxj.Map(vmap).SetValueForPath(acmt, "#attr.ios-CFBundleVersion.#text")
		err = m.SetValueForPath(vmap, "widget")
		check(err)

		// edit android-versionCode attr
		acmt, err = mxj.Map(vmap).ValueForPath("#attr.android-versionCode.#text")
		acmt = buildAndroidVersionCode(newVersion)
		mxj.Map(vmap).SetValueForPath(acmt, "#attr.android-versionCode.#text")
		err = m.SetValueForPath(vmap, "widget")
		check(err)

		b, err := m.XmlSeqIndent("", "  ")
		check(err)

		// Write header
		header := `<?xml version="1.0" encoding="utf-8"?>` + "\n"
		return []byte(header + string(b)), nil
	}

	return nil, nil
}
