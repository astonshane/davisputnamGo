package connector

import "testing"

func TestIsSimpleOr(t *testing.T) {

	cases := []struct {
		plaintext string
		want      bool
	}{
		{"A", true},
		{"~A", true},
		{"A^B", false},
		{"A^~B", false},
		{"A^~A^~B", false},
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
		got := Parse(c.plaintext).isSimpleOr()
		want := c.want
		if got != want {
			t.Errorf("Parse(%q).isSimpleOr(): %t != %t", c.plaintext, got, want)
		}
	}
}
