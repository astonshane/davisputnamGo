package connector

import "testing"

func TestIsCNF(t *testing.T) {

	cases := []struct {
		plaintext string
		want      bool
	}{
		{"A", true},
		{"~A", true},
		{"A^B", true},
		{"A^~B", true},
		{"A^~A^~B", true},
		{"AvB", true},
		{"Av~B", true},
		{"~Av~B", true},
		{"Av~Av~B", true},
		{"A->B", false},
		{"A->~B", false},
		{"A<->B", false},
		{"A<->~B", false},
	}
	for _, c := range cases {
		got := Parse(c.plaintext).isCNF()
		want := c.want
		if got != want {
			t.Errorf("Parse(%q).isCNF(): %t != %t", c.plaintext, got, want)
		}
	}
}
