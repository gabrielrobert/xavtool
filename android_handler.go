package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
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
	fileReader := bytes.NewReader(readFile(file.Path))
	for m, err := mxj.NewMapXmlSeqReader(fileReader); m != nil || err != io.EOF; m, err = mxj.NewMapXmlSeqReader(fileReader) {
		if err != nil {
			if err == mxj.NO_ROOT {
				continue
			} else {
				check(err)
			}
		}
		vmap := m["manifest"].(map[string]interface{})
		acmt, err := mxj.Map(vmap).ValueForPath("#attr.android:versionName.#text")
		acmt = newVersion
		mxj.Map(vmap).SetValueForPath(acmt, "#attr.android:versionName.#text")
		err = m.SetValueForPath(vmap, "manifest")
		if err != nil {
			fmt.Println("SetValueForPath err:", err)
			break
		}
		b, err := m.XmlSeqIndent("", "  ")
		check(err)
		saveFile(file.Path, b)
	}

	return nil
}
