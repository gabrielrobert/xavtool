package main

import (
	"testing"
)

func TestCurrent(t *testing.T) {
	current(nil)
	t.Errorf("credentials file mismatch")
}
