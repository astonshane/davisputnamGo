package clause

import (
	"davisputnam/literal"
	"testing"
)

func TestTautology(t *testing.T) {
	a, b, nb := literal.ConstructTestLiterals()

	clause := Clause{}
	clause.Append(a)
	clause.Append(b)

	clause2 := clause.Copy()
	clause2.Append(nb)

	cases := []struct {
		clause Clause
		want   bool
	}{
		{clause, false},
		{clause2, true},
	}

	for _, c := range cases {
		got := c.clause.Tautology()
		if got != c.want {
			t.Errorf("%q.Tautology() == %t, want %t", c.clause, got, c.want)
		}
	}
}
