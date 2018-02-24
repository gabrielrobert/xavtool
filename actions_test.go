package main

import (
	"testing"
)

func Test_ShouldReturnVersion(t *testing.T) {
	currentVersion := getCurrentVersion("test/Info.plist")
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}
}
