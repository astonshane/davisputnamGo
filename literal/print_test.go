package literal

import "testing"

func TestPrint(t *testing.T) {
	cases := []struct {
		inName    string
		inNegated bool
		want      string
	}{
		{"A", false, "A"},
		{"B", true, "~B"},
	}
	for _, c := range cases {
		l := Literal{Name: c.inName, Negated: c.inNegated}
		got := l.Print()
		if got != c.want {
			t.Errorf("Print(%q, %t) == %q, want %q", c.inName, c.inNegated, got, c.want)
		}
	}
}
