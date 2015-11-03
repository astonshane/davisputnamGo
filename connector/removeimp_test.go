package connector

import "testing"

func TestRemoveImp(t *testing.T) {

	cases := []struct {
		plaintext string
		want      string
	}{
		{"A->B", "{Or: [{Neg: {Literal: A}} {Literal: B}]}"},
		{"~A->B", "{Or: [{Literal: A} {Literal: B}]}"},
		{"A", "{Literal: A}"},
	}
	for _, c := range cases {
		got := Parse(c.plaintext).RemoveImp().String()
		if got != c.want {
			t.Errorf("Parse(%q).Negate().PropogateNegations(): %q != %q", c.plaintext, got, c.want)
		}
	}
}
