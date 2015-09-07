package literal

import "testing"

func TestCopy(t *testing.T) {
	cases := []struct {
		inName    string
		inNegated bool
		change    bool
		want      bool
	}{
		{"A", false, false, true},
		{"A", true, true, false},
	}
	for _, c := range cases {
		l := Literal{Name: c.inName, Negated: c.inNegated}
		copy := l.Copy()

		if c.change {
			copy.Name = "B"
		}

		if (l == copy) != c.want {
			t.Errorf("(%q == %q) == %t; want %t", l, copy, l == copy, c.want)
		}
	}
}
