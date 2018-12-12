package main

import (
	"os"
	"testing"
)

func Test_FindManifests(t *testing.T) {
	tables := []struct {
		file string
	}{
		{"Info.plist"},
		{"AndroidManifest.xml"},
		{"Package.appxmanifest"},
	}

	// scan current folder
	dir, err := os.Getwd()
	check(err)
	foundFiles, err := findManifests(dir, []packageHandler{iOSHandler{}, androidHandler{}, uwpHandler{}})

	if foundFiles == nil || len(foundFiles) == 0 {
		t.Errorf("at least one file should be found")
	}

	for _, table := range tables {
		if !containsStructFieldValue(foundFiles, "Path", table.file) {
			t.Errorf("test file info.plist has not been found, expected %v", "test/Info.plist")
		}
	}
}
func Test_FindManifests_ShouldIgnoreCordovaLibFolder(t *testing.T) {

	// scan current folder
	dir, err := os.Getwd()
	check(err)
	foundFiles, err := findManifests(dir+"\\test\\CordovaLib", []packageHandler{iOSHandler{}, androidHandler{}, uwpHandler{}})

	if len(foundFiles) == 1 {
		t.Errorf("should not be picked")
	}
}
