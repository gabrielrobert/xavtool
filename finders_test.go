package main

import (
	"testing"
)

func Test_ShouldFindiOSFile(t *testing.T) {
	foundFiles := findiOSFile("./")
	if foundFiles == nil || len(foundFiles) == 0 {
		t.Errorf("test file info.plist has not been found, expected %v", "test/Info.plist")
	}
}
