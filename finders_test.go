package main

import (
	"os"
	"testing"
)

func Test_FindManifests(t *testing.T) {
	// scan current folder
	dir, err := os.Getwd()
	check(err)

	foundFiles, err := findManifests(dir, []packageHandler{iOSHandler{}, androidHandler{}})
	if foundFiles == nil || len(foundFiles) == 0 {
		t.Errorf("at least one file should be found")
	}

	if !containsStructFieldValue(foundFiles, "Path", "Info.plist") {
		t.Errorf("test file info.plist has not been found, expected %v", "test/Info.plist")
	}

	if !containsStructFieldValue(foundFiles, "Path", "AndroidManifest.xml") {
		t.Errorf("test file androidmanifest.xml has not been found, expected %v", "test/AndroidManifest.xml")
	}

	if !containsStructFieldValue(foundFiles, "Path", "Package.appxmanifest") {
		t.Errorf("test file package.appxmanifest has not been found, expected %v", "test/Package.appxmanifest")
	}
}
