package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var uwpFilePath = "test/Package.appxmanifest"

func Test_uwpHandler_isPackage(t *testing.T) {
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
			handler := new(uwpHandler)
			if got := handler.isPackage(tt.args.filename); got != tt.want {
				t.Errorf("isPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uwpHandler_getPackageInfo(t *testing.T) {
	handler := new(uwpHandler)
	currentVersion, err := handler.getPackageInfo(uwpFilePath)
	require.NoError(t, err)
	assert.Equal(t, "1.0.1.0", currentVersion.Version)
	assert.Equal(t, "xavtool", currentVersion.Name)
	assert.Equal(t, "---", currentVersion.InternalVersion)
}

func Test_uwpHandler_changePackageVersion(t *testing.T) {
	handler := new(uwpHandler)
	currentVersion, err := handler.getPackageInfo(uwpFilePath)
	require.NoError(t, err)
	if currentVersion.Version != "1.0.1.0" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1.0")
	}

	handler.changePackageVersion(currentVersion, "1.0.2")
	currentVersion, err = handler.getPackageInfo(uwpFilePath)
	require.NoError(t, err)
	if currentVersion.Version != "1.0.2.0" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.2.0")
	}

	// some kind of rollback
	handler.changePackageVersion(currentVersion, "1.0.1")
}

func Test_uwpHandler_applyVersion(t *testing.T) {
	handler := new(uwpHandler)
	processedBytes, err := handler.applyVersion(uwpSeed, "1.0.2")
	require.NoError(t, err)
	xml, _ := handler.read(processedBytes)
	if xml.Identity.Version != "1.0.2.0" {
		t.Errorf("VersionName mismatch; expected %v, got %v", "1.0.2.0", xml.Identity.Version)
	}
}

func Test_uwpHandler_read(t *testing.T) {
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
		{"valid file", args{readFile("test/Package.appxmanifest")}, "1.0.1.0", false},
		{"valid bytes", args{uwpSeed}, "1.0.1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := new(uwpHandler)
			got, err := handler.read(tt.args.data)
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
