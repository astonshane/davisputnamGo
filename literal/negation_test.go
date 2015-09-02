package literal

import "testing"

func TestNegation(t *testing.T) {
	cases := []struct {
		inName    string
		inNegated bool
		want      string
	}{
		{"A", false, "~A"},
		{"B", true, "B"},
	}
	for _, c := range cases {
		l := Literal{Name: c.inName, Negated: c.inNegated}
		neg := l.Negation()
		got := neg.Print()
		if got != c.want {
			t.Errorf("Negation(%q) == %q, want %q", l.Print(), got, c.want)
		}
	}
}
