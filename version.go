package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
)

func incrementMajor(version string) string {
	parsedVersion := parse(version)
	newVersion := parsedVersion.IncMajor()
	return newVersion.String()
}

func incrementMinor(version string) string {
	parsedVersion := parse(version)
	newVersion := parsedVersion.IncMinor()
	return newVersion.String()
}

func incrementPatch(version string) string {
	parsedVersion := parse(version)
	newVersion := parsedVersion.IncPatch()
	return newVersion.String()
}

func parse(version string) *semver.Version {
	if !isVersion(version) {
		segments := strings.Split(version, ".")
		version = fmt.Sprintf("%v.%v.%v", segments[0], segments[1], segments[2])
	}

	parsedVersion, err := semver.NewVersion(version)
	check(err)
	return parsedVersion
}

func isVersion(version string) bool {
	_, err := semver.NewVersion(version)
	if err != nil {
		return false
	}
	return true
}

func buildAndroidVersionCode(version string) string {
	parsedVersion := parse(version)

	versionCode := parsedVersion.Major() * 1000000
	versionCode += parsedVersion.Minor() * 10000
	versionCode += parsedVersion.Patch() * 100

	if versionCode > 2000000000 {
		panic(fmt.Sprintf("Android versionCode cannot be greater than %v", 2000000000))
	}

	return strconv.FormatInt(versionCode, 10)
}

func buildUWPVersion(version string) string {
	return version + ".0"
}
