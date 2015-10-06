package connector

import "testing"

func TestIsLiteral(t *testing.T) {

	cases := []struct {
		plaintext string
		want      bool
	}{
		{"A", true},
		{"~A", true},
		{"A^B", false},
		{"AvB", false},
	}
	for _, c := range cases {
		got := Parse(c.plaintext).isLiteral()
		want := c.want
		if got != want {
			t.Errorf("Parse(%q).isLiteral(): %t != %t", c.plaintext, got, want)
		}
	}
}
