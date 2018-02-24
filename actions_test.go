package main

import (
	"testing"
)

func TestCurrent(t *testing.T) {
	currentVersion := getCurrentVersion()
	if currentVersion != "1.0.1" {
		t.Errorf("version mismatch; actual %v, expected %v", currentVersion, "1.0.1")
	}
}
