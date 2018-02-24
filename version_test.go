package main

import "testing"

func Test_parseVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"should parse strict minimal version", args{version: "0.0.0"}, "0.0.0"},
		{"should parse missing part version", args{version: "0.0"}, "0.0.0"},
		{"should parse missing part version", args{version: "0"}, "0.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseVersion(tt.args.version); got != tt.want {
				t.Errorf("parseVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
