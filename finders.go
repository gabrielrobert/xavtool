package main

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
)

type packageInfo struct {
	Name     string
	Version  string
	Path     string
	HasError bool
}

func findManifests(root string, handlers []packageHandler) ([]packageInfo, error) {
	var result error
	fileList := []packageInfo{}

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {

		if isIgnored(f) {
			return filepath.SkipDir
		}

		for _, handler := range handlers {
			if handler.isPackage(path) {
				pkg, err := handler.getPackageInfo(path)
				fileList = append(fileList, pkg)
				if err != nil {
					result = multierror.Append(result, err)
				}
			}
		}

		return nil
	})

	if err != nil {
		result = multierror.Append(result, err)
	}

	return fileList, result
}

func isIgnored(f os.FileInfo) bool {
	if f.IsDir() && stringInSlice(f.Name(), []string{"bin", "obj", ".git"}) {
		return true
	}
	return false
}
