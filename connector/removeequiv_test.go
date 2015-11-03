package connector

import "testing"

func TestRemoveEquiv(t *testing.T) {

	cases := []struct {
		plaintext string
		want      string
	}{
		{"A<->B", "{And: [{Or: [{Neg: {Literal: A}} {Literal: B}]} {Or: [{Neg: {Literal: B}} {Literal: A}]}]}"},
		{"~A<->B", "{And: [{Or: [{Literal: A} {Literal: B}]} {Or: [{Neg: {Literal: B}} {Neg: {Literal: A}}]}]}"},
		{"A", "{Literal: A}"},
	}
	for _, c := range cases {
		got := Parse(c.plaintext).RemoveEquiv().String()
		if got != c.want {
			t.Errorf("Parse(%q).RemoveEquiv(): %q != %q", c.plaintext, got, c.want)
		}
	}
}
