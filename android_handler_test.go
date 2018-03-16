package main

import (
	"testing"
)

func Test_isAndroidPackage(t *testing.T) {
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
			if got := isAndroidPackage(tt.args.filename); got != tt.want {
				t.Errorf("isAndroidPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAndroidPackageInfo(t *testing.T) {
	currentVersion := getAndroidPackageInfo(filePath)
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}
}

func Test_changeAndroidPackageVersion(t *testing.T) {
	currentVersion := getAndroidPackageInfo(filePath)
	if currentVersion.Version != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}

	changeAndroidPackageVersion(currentVersion, "1.0.2")
	currentVersion = getAndroidPackageInfo(filePath)
	if currentVersion.Version != "1.0.2" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.2")
	}

	// some kind of rollback
	changeAndroidPackageVersion(currentVersion, "1.0.1")
}

func Test_applyVersionToAndroidXML(t *testing.T) {
	processedBytes := applyVersionToAndroidXML(androidSeed, "1.0.2")
	xml := readAndroidData(processedBytes)
	if xml.VersionName != "1.0.2" {
		t.Errorf("VersionName mismatch; expected %v", "1.0.2")
	}
	if xml.Code != "102" {
		t.Errorf("code mismatch; expected %v", "102")
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
