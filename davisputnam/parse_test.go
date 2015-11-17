package main

import "testing"

func TestConstructs(t *testing.T) {
	cases := []struct {
		file string
		want string
	}{
		//should make each line in the input file (but the negation of the last line) a clause (or set of clauses)
		{"../inputs/test1.txt", "{{A}, {~A}}"},
		{"../inputs/test2.txt", "{{A}, {B}}"},
		{"../inputs/test3.txt", "{{A}, {B}, {~A, B}}"},
		{"../inputs/test4.txt", "{{A}, {~B}}"},
		{"../inputs/test5.txt", "{{A}, {~B}, {~A, B}}"},
	}
	for _, c := range cases {
		got := ConstructCS(c.file).String()
		if got != c.want {
			t.Errorf("ConstructCS(%q): %q != %q", c.file, got, c.want)
		}
	}
}
