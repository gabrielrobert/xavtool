package main

import (
	"testing"
)

func Test_incrementMajor(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1.0.0", args{"1.0.0"}, "2.0.0"},
		{"2.2.1", args{"2.2.1"}, "3.0.0"},
		{"0.0.0", args{"0.0.0"}, "1.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incrementMajor(tt.args.version); got != tt.want {
				t.Errorf("incrementMajor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_incrementMinor(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1.0.0", args{"1.0.0"}, "1.1.0"},
		{"2.2.1", args{"2.2.1"}, "2.3.0"},
		{"0.0.0", args{"0.0.0"}, "0.1.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incrementMinor(tt.args.version); got != tt.want {
				t.Errorf("incrementMinor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_incrementPatch(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1.0.0", args{"1.0.0"}, "1.0.1"},
		{"2.2.1", args{"2.2.1"}, "2.2.2"},
		{"0.0.0", args{"0.0.0"}, "0.0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incrementPatch(tt.args.version); got != tt.want {
				t.Errorf("incrementPatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"minimal version", args{"0.0.0"}, true},
		{"alpha value", args{"bleh"}, false},
		{"version with tag", args{"1.0.0-alpha"}, true},
		{"weird version", args{"1.aplpha"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVersion(tt.args.version); got != tt.want {
				t.Errorf("isVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildAndroidVersionCode(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1.99.1", args{"1.99.1"}, "1990100"},
		{"2.0.0", args{"2.0.0"}, "2000000"},
		{"1.0.1", args{"1.0.1"}, "1000100"},
		{"1.1.0", args{"1.1.0"}, "1010000"},
		{"0.0.0", args{"0.0.0"}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildAndroidVersionCode(tt.args.version); got != tt.want {
				t.Errorf("buildAndroidVersionCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildUWPVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1.99.1", args{"1.99.1"}, "1.99.1.0"},
		{"2.0.0", args{"2.0.0"}, "2.0.0.0"},
		{"1.0.1", args{"1.0.1"}, "1.0.1.0"},
		{"1.1.0", args{"1.1.0"}, "1.1.0.0"},
		{"0.0.0", args{"0.0.0"}, "0.0.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildUWPVersion(tt.args.version); got != tt.want {
				t.Errorf("buildUWPVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
