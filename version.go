package main

import (
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
