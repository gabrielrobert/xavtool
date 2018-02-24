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

	if !stringContainsInSlice("Info.plist", foundFiles) {
		t.Errorf("test file info.plist has not been found, expected %v", "test/Info.plist")
	}

	if !stringContainsInSlice("AndroidManifest.xml", foundFiles) {
		t.Errorf("test file AndroidManifest.xml has not been found, expected %v", "test/AndroidManifest.xml")
	}
}
