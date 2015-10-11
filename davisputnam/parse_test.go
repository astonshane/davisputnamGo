package main

import "testing"

func TestConstructs(t *testing.T) {
	cases := []struct {
		file string
		want string
	}{
		{"../inputs/test.txt", "{{A}}"},
		{"../inputs/test2.txt", "{{A}, {~B}}"},
		{"../inputs/test3.txt", "{{A}, {~B}, {~A, B}}"},
	}
	for _, c := range cases {
		got := ConstructCS(c.file).String()
		if got != c.want {
			t.Errorf("ConstructCS(%q): %q != %q", c.file, got, c.want)
		}
	}
}
