package main

import (
	"log"
	"os"
	"testing"
)

func TestFindManifests(t *testing.T) {
	// scan current folder
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	foundFiles := findManifests(dir)
	if foundFiles == nil || len(foundFiles) == 0 {
		t.Errorf("at least one file should be found")
	}

	if !containsStructFieldValue(foundFiles, "Path", "Info.plist") {
		t.Errorf("test file info.plist has not been found, expected %v", "test/Info.plist")
	}

	if !containsStructFieldValue(foundFiles, "Path", "AndroidManifest.xml") {
		t.Errorf("test file androidmanifest.xml has not been found, expected %v", "test/AndroidManifest.xml")
	}
}
