package connector

import "testing"

func TestParse(t *testing.T) {

	cases := []struct {
		plaintext string
		want      string
	}{
		{"A", "{Literal: A}"},
		{"B", "{Literal: B}"},
		{"~A", "{Neg: {Literal: A}}"},
		{"~B", "{Neg: {Literal: B}}"},
		{"A^B", "{And: [{Literal: A} {Literal: B}]}"},
		{"A^~B", "{And: [{Literal: A} {Neg: {Literal: B}}]}"},
		{"~A^~B", "{And: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{"A^~A^~B", "{And: [{Literal: A} {Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{"AvB", "{Or: [{Literal: A} {Literal: B}]}"},
		{"Av~B", "{Or: [{Literal: A} {Neg: {Literal: B}}]}"},
		{"~Av~B", "{Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{"Av~Av~B", "{Or: [{Literal: A} {Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{"A->B", "{Imp: [{Literal: A} {Literal: B}]}"},
		{"A->~B", "{Imp: [{Literal: A} {Neg: {Literal: B}}]}"},
		{"~A->~B", "{Imp: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{"A<->B", "{Equiv: [{Literal: A} {Literal: B}]}"},
		{"A<->~B", "{Equiv: [{Literal: A} {Neg: {Literal: B}}]}"},
		{"~A<->~B", "{Equiv: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
	}
	for _, c := range cases {
		got := Parse(c.plaintext).String()
		want := c.want
		if got != want {
			t.Errorf("Parse(%q): %q != %q", c.plaintext, got, want)
		}
	}
}
