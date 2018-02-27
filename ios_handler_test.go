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

func Test_changeiOSPackageVersion(t *testing.T) {
	currentVersion := getiOSPackageInfo("test/Info.plist")
	changeiOSPackageVersion(currentVersion, "1.0.2")
	currentVersion = getiOSPackageInfo("test/Info.plist")
	if currentVersion.Version != "1.0.2" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.2")
	}
}
