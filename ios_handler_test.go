package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_isPackage(t *testing.T) {
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
			handler := new(iOSHandler)
			if got := handler.isPackage(tt.args.filename); got != tt.want {
				t.Errorf("isiOsPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPackageInfo(t *testing.T) {
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
			handler := new(iOSHandler)
			got, _ := handler.getPackageInfo(tt.args.filename)
			if got.Version != tt.want.Version {
				t.Errorf("getPackageInfo.Version() = %v, want %v", got.Version, tt.want.Version)
			}
		})
	}
}

func Test_changePackageVersion(t *testing.T) {
	handler := new(iOSHandler)
	currentVersion, _ := handler.getPackageInfo("test/Info.plist")
	handler.changePackageVersion(currentVersion, "1.0.2")
	currentVersion, _ = handler.getPackageInfo("test/Info.plist")
	if currentVersion.Version != "1.0.2" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.2")
	}

	// some kind of rollback
	handler.changePackageVersion(currentVersion, "1.0.1")
}

func Test_applyVersion(t *testing.T) {
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
			handler := new(iOSHandler)
			got, err := handler.applyVersion(tt.args.data, tt.args.version)
			if tt.shouldError {
				require.Error(t, err)
				return
			}

			data, _ := handler.read(got)

			require.NoError(t, err)
			assert.Equal(t, tt.want, data["CFBundleVersion"])
		})
	}
}

func Test_read(t *testing.T) {
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
		{"missing properties", args{missingPropertiesPlist}, "", true},
		{"valid file", args{readFile("test/Info.plist")}, "1.0.1", false},
		{"valid bytes", args{iOSSeed}, "1.0.1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := new(iOSHandler)
			got, err := handler.read(tt.args.data)
			if tt.shouldError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got["CFBundleVersion"])
			assert.Equal(t, tt.want, got["CFBundleShortVersionString"])
		})
	}
}

var iOSSeed = []byte(`
	<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
	<plist version="1.0">
		<dict>
			<key>CFBundleDisplayName</key>
			<string>test</string>
			<key>CFBundleVersion</key>
			<string>1.0.1</string>
			<key>CFBundleShortVersionString</key>
			<string>1.0.1</string>
		</dict>
	</plist>
`)

var invalidPlist = []byte(`
	<?xml version="1.0" encoding="UTF-8"?>
	<plist ve/
`)

var missingPropertiesPlist = []byte(`
	<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
	<plist version="1.0">
		<dict>
		</dict>
	</plist>
`)
