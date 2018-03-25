package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"

	"github.com/clbanning/mxj"
)

type uwpIdentity struct {
	Version string `xml:"Version,attr"`
}

type uwpProperties struct {
	Name string `xml:"DisplayName"`
}

type uwpBundlerHeader struct {
	XMLName    xml.Name      `xml:"Package"`
	Identity   uwpIdentity   `xml:"Identity"`
	Properties uwpProperties `xml:"Properties"`
	Attrs      []xml.Attr    `xml:",attr"`
}

func isUWPPackage(filename string) bool {
	return strings.ToLower(getFilename(filename)) == "package.appxmanifest"
}

func getUWPPackageInfo(filePath string) packageInfo {
	byteValue := readFile(filePath)
	data, err := readUWPData(byteValue)

	if err != nil {
		return packageInfo{Path: filePath, HasError: true}
	}

	return packageInfo{
		Name:    data.Properties.Name,
		Version: data.Identity.Version,
		Path:    filePath,
	}
}

func readUWPData(data []byte) (*uwpBundlerHeader, error) {
	var header uwpBundlerHeader
	err := xml.Unmarshal(data, &header)
	return &header, err
}

func changeUWPPackageVersion(file packageInfo, newVersion string) error {
	fileBytes := readFile(file.Path)
	processedBytes := applyVersionToUWPXML(fileBytes, newVersion)
	saveFile(file.Path, processedBytes)
	return nil
}

func applyVersionToUWPXML(data []byte, newVersion string) []byte {
	fileReader := bytes.NewReader(data)
	for m, err := mxj.NewMapXmlSeqReader(fileReader); m != nil || err != io.EOF; m, err = mxj.NewMapXmlSeqReader(fileReader) {
		if err != nil {
			if err == mxj.NO_ROOT {
				continue
			} else {
				check(err)
			}
		}
		temp := m["Package"].(map[string]interface{})
		vmap := temp["Identity"].(map[string]interface{})

		// edit Version attr
		acmt, err := mxj.Map(vmap).ValueForPath("#attr.Version.#text")
		acmt = buildUWPVersion(version)
		mxj.Map(vmap).SetValueForPath(acmt, "#attr.Version.#text")
		err = m.SetValueForPath(vmap, "Package.Identity")
		check(err)

		b, err := m.XmlSeqIndent("", "  ")
		check(err)

		// Write header
		header := `<?xml version="1.0" encoding="utf-8"?>` + "\n"
		return []byte(header + string(b))
	}

	return nil
}
