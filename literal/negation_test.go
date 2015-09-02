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
		//test that the negation is what we want
		if got != c.want {
			t.Errorf("Negation(%q) == %q, want %q", l.Print(), got, c.want)
		}
		//test that the negation of the negation is what we statted with
		original := neg.Negation()
		if l.Print() != original.Print() {
			t.Errorf("Negation(Negation(%q)) == %q, want %q", l.Print(), original.Print(), l.Print())
		}
	}
}
