package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func findiOSFile(root string) []string {
	fileList := []string{}

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, "Info.plist") {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return fileList
}
