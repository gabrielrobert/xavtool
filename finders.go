package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ignoreFolders = []string{"bin", "obj", ".git"}

func findiOSFile(root string) []string {
	fileList := []string{}

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {

		// ignore specific folders
		if toBeIgnored(f) {
			return filepath.SkipDir
		}

		if isiOSFile(f) {
			fileList = append(fileList, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return fileList
}

func toBeIgnored(f os.FileInfo) bool {
	if f.IsDir() && stringInSlice(f.Name(), ignoreFolders) {
		return true
	}
	return false
}

func isiOSFile(f os.FileInfo) bool {
	if !f.IsDir() && strings.ToLower(f.Name()) == "info.plist" {
		return true
	}
	return false
}
