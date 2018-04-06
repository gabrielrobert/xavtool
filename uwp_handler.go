package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/clbanning/mxj"
)

type uwpHandler struct {
}

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

func (h uwpHandler) isPackage(filename string) bool {
	return strings.ToLower(getFilename(filename)) == "package.appxmanifest"
}

func (h uwpHandler) getPackageInfo(filePath string) (packageInfo, error) {
	byteValue := readFile(filePath)
	data, err := h.read(byteValue)

	if err != nil {
		return packageInfo{Path: filePath, HasError: true}, err
	}

	return packageInfo{
		Name:            data.Properties.Name,
		Version:         data.Identity.Version,
		InternalVersion: "---",
		Path:            filePath,
	}, nil
}

func (h uwpHandler) read(data []byte) (*uwpBundlerHeader, error) {
	var header uwpBundlerHeader
	err := xml.Unmarshal(data, &header)
	return &header, err
}

func (h uwpHandler) changePackageVersion(file packageInfo, newVersion string) error {
	fileBytes := readFile(file.Path)
	processedBytes, err := h.applyVersion(fileBytes, newVersion)
	if err != nil {
		return fmt.Errorf("Invalid xml file: %v", file.Path)
	}
	saveFile(file.Path, processedBytes)
	return nil
}

func (h uwpHandler) applyVersion(byteValue []byte, newVersion string) ([]byte, error) {
	fileReader := bytes.NewReader(byteValue)
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
		acmt = buildUWPVersion(newVersion)
		mxj.Map(vmap).SetValueForPath(acmt, "#attr.Version.#text")
		err = m.SetValueForPath(vmap, "Package.Identity")
		check(err)

		b, err := m.XmlSeqIndent("", "  ")
		check(err)

		// Write header
		header := `<?xml version="1.0" encoding="utf-8"?>` + "\n"
		return []byte(header + string(b)), nil
	}

	return nil, nil
}
