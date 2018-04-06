package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_androidHandler_isPackage(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"filename", args{"AndroidmanifesT.xmL"}, true},
		{"filename with path and \\", args{"c:/dev/androidmanifest.xml"}, true},
		{"filename with path and /", args{"c:\\dev\\androidmanifest.xml"}, true},
		{"filename with path", args{"c:\\dev/androidmanifest.xml"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := new(androidHandler)
			if got := handler.isPackage(tt.args.filename); got != tt.want {
				t.Errorf("isAndroidPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_androidHandler_getPackageInfo(t *testing.T) {
	handler := new(androidHandler)
	currentVersion, err := handler.getPackageInfo(filePath)
	require.NoError(t, err)
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}
}

func Test_androidHandler_changePackageVersion(t *testing.T) {
	handler := new(androidHandler)
	currentVersion, err := handler.getPackageInfo(filePath)
	require.NoError(t, err)
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}

	handler.changePackageVersion(currentVersion, "1.0.2")
	currentVersion, err = handler.getPackageInfo(filePath)
	require.NoError(t, err)
	if currentVersion.Version != "1.0.2" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.2")
	}

	// some kind of rollback
	handler.changePackageVersion(currentVersion, "1.0.1")
}

func Test_androidHandler_applyVersion(t *testing.T) {
	handler := new(androidHandler)
	processedBytes, err := handler.applyVersion(androidSeed, "1.0.2")
	require.NoError(t, err)
	xml, _ := handler.read(processedBytes)
	if xml.VersionName != "1.0.2" {
		t.Errorf("VersionName mismatch; expected %v", "1.0.2")
	}
	if xml.Code != "1000200" {
		t.Errorf("code mismatch; actual %v, expected %v", xml.Code, "1000200")
	}
}

func Test_androidHandler_read(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name        string
		args        args
		want        string
		shouldError bool
	}{
		{"invalid bytes", args{invalidAndroidSeed}, "", true},
		{"valid file", args{readFile("test/AndroidManifest.xml")}, "1.0.1", false},
		{"valid bytes", args{androidSeed}, "1.0.1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := new(androidHandler)
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

var filePath = "test/AndroidManifest.xml"

var androidSeed = []byte(`
	<?xml version="1.0" encoding="utf-8"?>
	<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    package="com.example.xavtool"
    android:versionCode="1"
    android:versionName="1.0.1" >
 	   <permission android:name="android"></permission>
 
		<application
			android:allowBackup="true"
			android:icon="@drawable/ic_launcher"
			android:label="@string/app_name"
			android:theme="@style/Theme.Sample" >
			<activity
				android:name="com.example.xavtool.MainActivity"
				android:label="@string/app_name"
				android:launchMode="singleTop">
				<meta-data
					android:name="android.app.searchable"
					android:resource="@xml/searchable" />
				<intent-filter>
					<action android:name="android.intent.action.SEARCH" />
				</intent-filter>
				<intent-filter>
					<action android:name="android.intent.action.MAIN" />
					<category android:name="android.intent.category.LAUNCHER" />
				</intent-filter>
			</activity>
		</application>
	</manifest>
`)

var invalidAndroidSeed = []byte(`
	<?xml version="1.0" encoding="utf-8"?>
	<man
`)
