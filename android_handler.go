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

type androidHandler struct {
}

type androidBundlerHeader struct {
	XMLName xml.Name `xml:"manifest"`
	Name    string   `xml:"package,attr"`

	// should be less than 2100000000
	Code string `xml:"versionCode,attr"`

	VersionName string     `xml:"versionName,attr"`
	Attrs       []xml.Attr `xml:",attr"`
}

func (h androidHandler) isPackage(filename string) bool {
	return strings.ToLower(getFilename(filename)) == "androidmanifest.xml"
}

func (h androidHandler) getPackageInfo(filePath string) (packageInfo, error) {
	byteValue := readFile(filePath)
	data, err := h.read(byteValue)

	if err != nil {
		return packageInfo{Path: filePath, HasError: true}, err
	}

	return packageInfo{
		Name:    data.Name,
		Version: data.VersionName,
		Path:    filePath,
	}, nil
}

func (h androidHandler) read(data []byte) (*androidBundlerHeader, error) {
	var header androidBundlerHeader
	err := xml.Unmarshal(data, &header)
	return &header, err
}

func (h androidHandler) changePackageVersion(file packageInfo, newVersion string) error {
	fileBytes := readFile(file.Path)
	processedBytes, err := h.applyVersion(fileBytes, newVersion)
	if err != nil {
		return fmt.Errorf("Invalid xml file: %v", file.Path)
	}
	saveFile(file.Path, processedBytes)
	return nil
}

func (h androidHandler) applyVersion(byteValue []byte, newVersion string) ([]byte, error) {
	fileReader := bytes.NewReader(byteValue)
	for m, err := mxj.NewMapXmlSeqReader(fileReader); m != nil || err != io.EOF; m, err = mxj.NewMapXmlSeqReader(fileReader) {
		if err != nil {
			if err == mxj.NO_ROOT {
				continue
			} else {
				return nil, errors.New("Invalid xml")
			}
		}
		vmap := m["manifest"].(map[string]interface{})

		// edit versionName attr
		acmt, err := mxj.Map(vmap).ValueForPath("#attr.android:versionName.#text")
		acmt = newVersion
		mxj.Map(vmap).SetValueForPath(acmt, "#attr.android:versionName.#text")
		err = m.SetValueForPath(vmap, "manifest")

		// edit versionCode attr
		acmt, err = mxj.Map(vmap).ValueForPath("#attr.android:versionCode.#text")
		acmt = buildAndroidVersionCode(newVersion)
		mxj.Map(vmap).SetValueForPath(acmt, "#attr.android:versionCode.#text")
		err = m.SetValueForPath(vmap, "manifest")
		check(err)

		b, err := m.XmlSeqIndent("", "  ")
		check(err)

		// Write header
		header := `<?xml version="1.0" encoding="utf-8"?>` + "\n"
		return []byte(header + string(b)), nil
	}

	return nil, nil
}
