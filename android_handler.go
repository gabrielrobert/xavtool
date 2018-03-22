package main

import (
	"bytes"
	"encoding/xml"
	"io"
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
	return strings.ToLower(getFilename(filename)) == "androidmanifest.xml"
}

func getAndroidPackageInfo(filePath string) packageInfo {
	byteValue := readFile(filePath)
	data, err := readAndroidData(byteValue)

	if err != nil {
		return packageInfo{Path: filePath, HasError: true}
	}

	return packageInfo{
		Name:    data.Name,
		Version: data.VersionName,
		Path:    filePath,
	}
}

func readAndroidData(data []byte) (*androidBundlerHeader, error) {
	var header androidBundlerHeader
	err := xml.Unmarshal(data, &header)
	return &header, err
}

func changeAndroidPackageVersion(file packageInfo, newVersion string) error {
	fileBytes := readFile(file.Path)
	processedBytes := applyVersionToAndroidXML(fileBytes, newVersion)
	saveFile(file.Path, processedBytes)
	return nil
}

func applyVersionToAndroidXML(data []byte, newVersion string) []byte {
	fileReader := bytes.NewReader(data)
	for m, err := mxj.NewMapXmlSeqReader(fileReader); m != nil || err != io.EOF; m, err = mxj.NewMapXmlSeqReader(fileReader) {
		if err != nil {
			if err == mxj.NO_ROOT {
				continue
			} else {
				check(err)
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
		return []byte(header + string(b))
	}

	return nil
}
