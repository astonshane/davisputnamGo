package connector

import "testing"

func TestToCNF(t *testing.T) {

	cases := []struct {
		plaintext string
		want      string
	}{
		{"A", "{Literal: A}"},
		{"~A", "{Neg: {Literal: A}}"},
		{"A^B", "{And: [{Literal: A} {Literal: B}]}"},
		{"A^~B", "{And: [{Literal: A} {Neg: {Literal: B}}]}"},
		{"A^~A^~B", "{And: [{Literal: A} {Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{"AvB", "{Or: [{Literal: A} {Literal: B}]}"},
		{"Av~B", "{Or: [{Literal: A} {Neg: {Literal: B}}]}"},
		{"~Av~B", "{Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{"Av~Av~B", "{Or: [{Literal: A} {Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		//{"A->B", "{Or: [{Neg: {Literal: A}} {Literal: B}]}"},
		//{"A->~B", "{Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		//{"A<->B", "{And: [{Or: [{Neg: {Literal: A}} {Literal: B}]} {Or: [{Literal: A} {Neg: {Literal: B}} ]}]}"},
		//{"A<->~B", "{And: [{Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]} {Or: [{Literal: A} {Literal: B}]}]}"},
	}
	for _, c := range cases {
		got := Parse(c.plaintext).ToCNF().String()
		if got != c.want {
			t.Errorf("Parse(%q).isCNF(): %q != %q", c.plaintext, got, c.want)
		}
	}
}
