package main

import (
	"log"

	"github.com/Masterminds/semver"
)

func parseVersion(version string) string {
	parsedVersion, err := semver.NewVersion(version)
	if err != nil {
		log.Fatal(err)
	}
	return parsedVersion.String()
}
