package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var uwpFilePath = "test/Package.appxmanifest"

func Test_isUWPPackage(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"filename", args{"PackAge.AppxmanifesT"}, true},
		{"filename with path and \\", args{"c:/dev/package.appxmanifest"}, true},
		{"filename with path and /", args{"c:\\dev\\package.appxmanifest"}, true},
		{"filename with path", args{"c:\\dev/package.appxmanifest"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUWPPackage(tt.args.filename); got != tt.want {
				t.Errorf("isUWPPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUWPPackageInfo(t *testing.T) {
	currentVersion := getUWPPackageInfo(uwpFilePath)
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion.Version, "1.0.1")
	}
	if currentVersion.Name != "xavtool" {
		t.Errorf("name mismatch; actual %v, expected %v", currentVersion.Name, "xavtool")
	}
}

func Test_changeUWPPackageVersion(t *testing.T) {
	currentVersion := getUWPPackageInfo(uwpFilePath)
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}

	changeUWPPackageVersion(currentVersion, "1.0.2")
	currentVersion = getUWPPackageInfo(uwpFilePath)
	if currentVersion.Version != "1.0.2" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.2")
	}

	// some kind of rollback
	changeUWPPackageVersion(currentVersion, "1.0.1")
}

func Test_applyVersionToUWPXML(t *testing.T) {
	processedBytes := applyVersionToUWPXML(uwpSeed, "1.0.2")
	xml, _ := readUWPData(processedBytes)
	if xml.Identity.Version != "1.0.2" {
		t.Errorf("VersionName mismatch; expected %v, got %v", "1.0.2", xml.Identity.Version)
	}
}

func Test_readUWPData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name        string
		args        args
		want        string
		shouldError bool
	}{
		{"invalid bytes", args{invalidUWPSeed}, "", true},
		{"valid file", args{readFile("test/Package.appxmanifest")}, "1.0.1", false},
		{"valid bytes", args{uwpSeed}, "1.0.1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readUWPData(tt.args.data)
			if tt.shouldError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got.Identity.Version)
		})
	}
}

var uwpSeed = []byte(`
	<?xml version="1.0" encoding="utf-8"?>
	<Package xmlns="http://schemas.microsoft.com/appx/manifest/foundation/windows10" xmlns:mp="http://schemas.microsoft.com/appx/2014/phone/manifest" xmlns:uap="http://schemas.microsoft.com/appx/manifest/uap/windows10" IgnorableNamespaces="uap mp">
  		<Identity Name="95748d56-342b-4dae-93f5-aeda0587a1c0" Publisher="CN=gabrielrobert" Version="1.0.1" />
  </Package>
`)

var invalidUWPSeed = []byte(`
	<?xml version="1.0" encoding="utf-8"?>
	<Pack/
`)
