package main

import (
	"testing"
)

func Test_getiOSPackageInfo(t *testing.T) {
	currentVersion := getiOSPackageInfo("test/Info.plist")
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}
}
