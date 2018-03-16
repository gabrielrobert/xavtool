package main

import (
	"testing"
)

func Test_isiOsPackage(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"filename", args{"Info.plisT"}, true},
		{"filename with path and \\", args{"c:/dev/info.plist"}, true},
		{"filename with path and /", args{"c:\\dev\\info.plist"}, true},
		{"filename with path", args{"c:\\dev/info.plist"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isiOsPackage(tt.args.filename); got != tt.want {
				t.Errorf("isiOsPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

	// some kind of rollback
	changeiOSPackageVersion(currentVersion, "1.0.1")
}

func Test_applyVersionToiOSPlist(t *testing.T) {
	processedBytes := applyVersionToiOSPlist(iOSSeed, "1.0.2")
	iOSDetails := readiOSData(processedBytes)
	if iOSDetails["CFBundleVersion"] != "1.0.2" {
		t.Errorf("VersionName mismatch; expected %v", "1.0.2")
	}
}

var iOSSeed = []byte(`
	<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
	<plist version="1.0">
		<dict>
			<key>CFBundleInfoDictionaryVersion</key>
			<string>6.0</string>
			<key>band-size</key>
			<integer>8388608</integer>
			<key>bundle-backingstore-version</key>
			<integer>1</integer>
			<key>diskimage-bundle-type</key>
			<string>com.apple.diskimage.sparsebundle</string>
			<key>size</key>
			<integer>4398046511104</integer>
		</dict>
	</plist>
`)
