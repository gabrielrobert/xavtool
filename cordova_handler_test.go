package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_cordovaHandler_isPackage(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"filename", args{"config.xmL"}, true},
		{"filename with path and \\", args{"c:/dev/config.xml"}, true},
		{"filename with path and /", args{"c:\\dev\\config.xml"}, true},
		{"filename with path", args{"c:\\dev/config.xml"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := new(cordovaHandler)
			if got := handler.isPackage(tt.args.filename); got != tt.want {
				t.Errorf("isCordovaPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cordovaHandler_getPackageInfo(t *testing.T) {
	handler := new(cordovaHandler)
	currentVersion, err := handler.getPackageInfo("test/config.xml")
	require.NoError(t, err)
	assert.Equal(t, "com.example.xavtool", currentVersion.Name)
	assert.Equal(t, "1.0.1", currentVersion.Version)
	assert.Equal(t, "1000100", currentVersion.InternalVersion)
}

func Test_cordovaHandler_changePackageVersion(t *testing.T) {
	handler := new(cordovaHandler)
	currentVersion, err := handler.getPackageInfo("test/config.xml")
	require.NoError(t, err)
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}

	handler.changePackageVersion(currentVersion, "1.0.2")
	currentVersion, err = handler.getPackageInfo("test/config.xml")
	require.NoError(t, err)
	if currentVersion.Version != "1.0.2" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.2")
	}

	// some kind of rollback
	handler.changePackageVersion(currentVersion, "1.0.1")
}

func Test_cordovaHandler_applyVersion(t *testing.T) {
	handler := new(cordovaHandler)
	processedBytes, err := handler.applyVersion(cordovaSeed, "1.0.2")
	require.NoError(t, err)
	xml, _ := handler.read(processedBytes)
	if xml.VersionName != "1.0.2" {
		t.Errorf("VersionName mismatch; expected %v", "1.0.2")
	}
	if xml.Code != "1000200" {
		t.Errorf("code mismatch; actual %v, expected %v", xml.Code, "1000200")
	}
}

func Test_cordovaHandler_read(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name        string
		args        args
		want        string
		shouldError bool
	}{
		{"invalid bytes", args{invalidCordovaSeed}, "", true},
		{"valid file", args{readFile("test/config.xml")}, "1.0.1", false},
		{"valid bytes", args{cordovaSeed}, "1.0.1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := new(cordovaHandler)
			got, err := handler.read(tt.args.data)
			if tt.shouldError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got.VersionName)
		})
	}
}

var cordovaSeed = []byte(`
  <?xml version="1.0" encoding="utf-8"?>
  <widget android-versionCode="1000100" id="com.example.xavtool" ios-CFBundleVersion="1.0.1" version="1.0.1" xmlns="http://www.w3.org/ns/widgets" xmlns:android="http://schemas.android.com/apk/res/android" xmlns:cdv="http://cordova.apache.org/ns/1.0">
    <name>xavtool</name>
    <content src="index.html"/>
    <access origin="*"/>
  </widget>
`)

var invalidCordovaSeed = []byte(`
	<?xml version="1.0" encoding="utf-8"?>
	<man
`)
