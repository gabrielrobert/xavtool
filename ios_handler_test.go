package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want packageInfo
	}{
		{"normal file", args{"test/Info.plist"}, packageInfo{Version: "1.0.1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getiOSPackageInfo(tt.args.filename); got.Version != tt.want.Version {
				t.Errorf("getiOSPackageInfo.Version() = %v, want %v", got.Version, tt.want.Version)
			}
		})
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
	type args struct {
		data    []byte
		version string
	}
	tests := []struct {
		name        string
		args        args
		want        string
		shouldError bool
	}{
		{"invalid bytes", args{invalidPlist, "1.0.2"}, "", true},
		{"valid file", args{readFile("test/Info.plist"), "1.0.2"}, "1.0.2", false},
		{"valid bytes", args{iOSSeed, "1.0.2"}, "1.0.2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyVersionToiOSPlist(tt.args.data, tt.args.version)
			if tt.shouldError {
				require.Error(t, err)
				return
			}

			data, _ := readiOSData(got)

			require.NoError(t, err)
			assert.Equal(t, tt.want, data["CFBundleVersion"])
		})
	}
}

func Test_readiOSData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name        string
		args        args
		want        string
		shouldError bool
	}{
		{"invalid bytes", args{invalidPlist}, "", true},
		{"valid file", args{readFile("test/Info.plist")}, "1.0.1", false},
		{"valid bytes", args{iOSSeed}, "1.0.1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readiOSData(tt.args.data)
			if tt.shouldError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got["CFBundleVersion"])
		})
	}
}

var iOSSeed = []byte(`
	<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
	<plist version="1.0">
		<dict>
			<key>CFBundleInfoDictionaryVersion</key>
			<string>6.0</string>
			<key>CFBundleVersion</key>
			<string>1.0.1</string>
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

var invalidPlist = []byte(`
	<?xml version="1.0" encoding="UTF-8"?>
	<plist ve/
`)
