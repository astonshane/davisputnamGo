package literal

import "testing"

func TestEquality(t *testing.T) {
	cases := []struct {
		aName    string
		aNegated bool
		bName    string
		bNegated bool
		want     bool
	}{
		{"A", false, "A", false, true},
		{"A", false, "A", true, false},
		{"A", false, "B", false, false},
		{"A", false, "B", true, false},
	}
	for _, c := range cases {
		a := Literal{c.aName, c.aNegated}
		b := Literal{c.bName, c.bNegated}

		if (a == b) != c.want {
			t.Errorf("Equals(%q, %q) == %t, want %t", a.String(), b.String(), a == b, c.want)
		}
	}
}
