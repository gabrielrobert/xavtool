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
