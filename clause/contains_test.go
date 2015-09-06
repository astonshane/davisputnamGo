package clause

import (
	"davisputnam/literal"
	"testing"
)

func TestContains(t *testing.T) {
	a, b, nb := literal.ConstructTestLiterals()

	clause := Clause{}
	clause.Append(a)
	clause.Append(b)

	cases := []struct {
		lit  literal.Literal
		want int
	}{
		{a, 0},
		{b, 1},
		{nb, -1},
	}

	for _, c := range cases {
		got := clause.Contains(c.lit)
		if got != c.want {
			t.Errorf("%q.Contains(%q) == %q, want %q", clause, c.lit, got, c.want)
		}
	}
}
