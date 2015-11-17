package main

import "testing"

func TestParseSat(t *testing.T) {
	cases := []struct {
		file string
		want bool
	}{
		//should make each line in the input file (but the negation of the last line) a clause (or set of clauses)
		{"../inputs/test1.txt", false},
		{"../inputs/test2.txt", true},
		{"../inputs/test3.txt", true},
		{"../inputs/test4.txt", true},
		{"../inputs/test5.txt", false},
	}
	for _, c := range cases {
		CS := ConstructCS(c.file)
		got := Satisfiable(CS)
		if got != c.want {
			t.Errorf("Satisfiable(%q): %t != %t", CS, got, c.want)
		}
	}
}
