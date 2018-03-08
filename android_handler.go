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

// Uncomment the print statements to details of the process here.
func copyCmts(m mxj.Map, path string) error {
	// get the array of Items entries for the 'path'
	vals, err := m.ValuesForPath(path)
	if err != nil {
		return fmt.Errorf("ValuesForPath err: %s", err.Error())
	} else if len(vals) == 0 {
		return fmt.Errorf("no vals for path: %s", path)
	}
	// process each Items entry
	for _, v := range vals {
		vm, ok := v.(map[string]interface{})
		if !ok {
			return fmt.Errorf("assertion failed")
		}
		// get the Comment list
		c, ok := vm["#attr.android:versionCode.#text"]
		if !ok { // --> no Items.Comment elements
			continue
		}
		// Don't assume that Comment is an array.
		// There may be just one value, in which case it will decode as map[string]interface{}.
		switch c.(type) {
		case map[string]interface{}:
			c = []interface{}{c}
		}
		cmt := c.([]interface{})
		// get the Request list
		r, ok := vm["Request"]
		if !ok { // --> no Items.Request elements
			continue
		}
		// Don't assume the Request is an array.
		// There may be just one value, in which case it will decode as map[string]interface{}.
		switch r.(type) {
		case map[string]interface{}:
			r = []interface{}{r}
		}
		req := r.([]interface{})

		// fmt.Println("Comment:", cmt)
		// fmt.Println("Request:", req)

		// Comment elements with #seq==n are followed by Request element with #seq==n+1.
		// For each Comment.#seq==n extract the CommentText attribute value and use it to
		// set the ReportingName attribute value in Request.#seq==n+1.
		for _, v := range cmt {
			vmap := v.(map[string]interface{})
			seq := vmap["#seq"].(int) // type is int
			// extract CommentText attr from "#attr"
			acmt, _ := mxj.Map(vmap).ValueForPath("#attr.CommentText.#text")
			if acmt == "" {
				fmt.Println("no CommentText value in Comment attributes")
			}
			// fmt.Println(seq, acmt)
			// find the request with the #seq==seq+1 value
			var r map[string]interface{}
			for _, vv := range req {
				rt := vv.(map[string]interface{})
				if rt["#seq"].(int) == seq+1 {
					r = rt
					break
				}
			}
			if r == nil { // no Request with #seq==seq+1
				continue
			}
			if err := mxj.Map(r).SetValueForPath(acmt, "#attr.ReportingName.#text"); err != nil {
				fmt.Println("SetValueForPath err:", err)
				break
			}
		}
	}
	return nil
}
