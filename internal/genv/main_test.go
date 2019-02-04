package main

import "testing"

func TestVersionNoPatch(t *testing.T) {
	data := []struct {
		in  string
		out string
	}{
		{"go1.11.4", "devel/release.html#go1.11.minor"},
		{"go1.12", "go1.12"},
		{"go1.12beta1", "go1.12"},
		{"go1.12rc2", "go1.12"},
	}
	for _, item := range data {
		if out := versionNoPatch(item.in); out != item.out {
			t.Errorf("versionNoPatch(%q) = %q; want %q", item.in, out, item.out)
		}
	}
}
