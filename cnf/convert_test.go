package cnf

import (
	"testing"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		plaintext string
		want      string
	}{
		{"A", "{{A}}"},
		{"~A", "{{~A}}"},
		{"A v B", "{{A, B}}"},
		{"A v B v C", "{{A, B, C}}"},
		{"~A v B", "{{A, B}}"},
		{"A v B v ~C", "{{A, B, C}}"},
	}

	for _, c := range cases {
		if Convert(c.plaintext).String() != c.want {
			t.Errorf("Convert(%q) == %q, not %q", c.plaintext, Convert(c.plaintext), c.want)
		}
	}
}
