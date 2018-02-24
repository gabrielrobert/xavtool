package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func Test_ShouldFindiOSFile(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	foundFiles := findiOSFile(dir)
	if foundFiles == nil || len(foundFiles) == 0 {
		t.Errorf("test file info.plist has not been found, expected %v", "test/Info.plist")
	}
}
