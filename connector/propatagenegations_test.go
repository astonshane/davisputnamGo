package connector

import "testing"

func TestPropogateNegations(t *testing.T) {

	cases := []struct {
		plaintext string
		want      string
	}{
		{"A", "{Neg: {Literal: A}}"},
		{"~A", "{Literal: A}"},
		{"A^B", "{Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{"A^~B", "{Or: [{Neg: {Literal: A}} {Literal: B}]}"},
		{"A^~A^~B", "{Or: [{Neg: {Literal: A}} {Literal: A} {Literal: B}]}"},
		{"AvB", "{And: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{"Av~B", "{And: [{Neg: {Literal: A}} {Literal: B}]}"},
		{"~Av~B", "{And: [{Literal: A} {Literal: B}]}"},
		{"Av~Av~B", "{And: [{Neg: {Literal: A}} {Literal: A} {Literal: B}]}"},
		//{"A->B", "{Neg: {Or: [{Neg: {Literal: A}} {Literal: B}]}}"},
		//{"A->~B", "{Neg: {Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}}"},
		//{"A<->B", "{Neg: {And: [{Or: [{Neg: {Literal: A}} {Literal: B}]} {Or: [{Literal: A} {Neg: {Literal: B}} ]}]}}"},
		//{"A<->~B", "{Neg: {And: [{Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]} {Or: [{Literal: A} {Literal: B}]}]}}"},

	}
	for _, c := range cases {
		got := Parse(c.plaintext).Negate().PropagateNegations().String()
		if got != c.want {
			t.Errorf("Parse(%q).Negate().PropogateNegations(): %q != %q", c.plaintext, got, c.want)
		}
	}
}
