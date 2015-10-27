package connector

import "testing"

func TestNegate(t *testing.T) {

	cases := []struct {
		plaintext string
		want      string
	}{
		{"A", "{Neg: {Literal: A}}"},
		{"~A", "{Neg: {Neg: {Literal: A}}}"},
		{"A^B", "{Neg: {And: [{Literal: A} {Literal: B}]}}"},
		{"A^~B", "{Neg: {And: [{Literal: A} {Neg: {Literal: B}}]}}"},
		{"A^~A^~B", "{Neg: {And: [{Literal: A} {Neg: {Literal: A}} {Neg: {Literal: B}}]}}"},
		{"AvB", "{Neg: {Or: [{Literal: A} {Literal: B}]}}"},
		{"Av~B", "{Neg: {Or: [{Literal: A} {Neg: {Literal: B}}]}}"},
		{"~Av~B", "{Neg: {Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}}"},
		{"Av~Av~B", "{Neg: {Or: [{Literal: A} {Neg: {Literal: A}} {Neg: {Literal: B}}]}}"},
		//{"A->B", "{Neg: {Or: [{Neg: {Literal: A}} {Literal: B}]}}"},
		//{"A->~B", "{Neg: {Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}}"},
		//{"A<->B", "{Neg: {And: [{Or: [{Neg: {Literal: A}} {Literal: B}]} {Or: [{Literal: A} {Neg: {Literal: B}} ]}]}}"},
		//{"A<->~B", "{Neg: {And: [{Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]} {Or: [{Literal: A} {Literal: B}]}]}}"},
	}
	for _, c := range cases {
		got := Parse(c.plaintext).Negate().String()
		if got != c.want {
			t.Errorf("Parse(%q).Negate(): %q != %q", c.plaintext, got, c.want)
		}
	}
}
