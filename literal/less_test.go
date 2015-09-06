package literal

import "testing"

func TestLess(t *testing.T) {
	cases := []struct {
		aName    string
		aNegated bool
		bName    string
		bNegated bool
		want     bool
	}{
		{"A", false, "A", false, false},
		{"A", false, "A", true, true},
		{"A", true, "A", false, false},
		{"A", false, "B", false, true},
		{"A", false, "B", true, true},
	}
	for _, c := range cases {
		a := Literal{c.aName, c.aNegated}
		b := Literal{c.bName, c.bNegated}

		if Less(a, b) != c.want {
			t.Errorf("Less(%q, %q) == %t, want %t", a.String(), b.String(), Less(a, b), c.want)
		}
	}
}
