package main

import (
	"testing"
)

func Test_getAndroidPackageInfo(t *testing.T) {
	currentVersion := getAndroidPackageInfo("test/AndroidManifest.xml")
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}
}

func Test_changeAndroidPackageVersion(t *testing.T) {
	currentVersion := getAndroidPackageInfo("test/AndroidManifest.xml")
	changeAndroidPackageVersion(currentVersion, "1.0.2")
	currentVersion = getAndroidPackageInfo("test/AndroidManifest.xml")
	if currentVersion.Version != "1.0.2" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.2")
	}
}
