package clause

import (
	"testing"
)

func TestSubsumption(t *testing.T) {
	empty, one, two, three := ConstructTestClauses()
	_, _, four, five := ConstructMoreTestClauses()

	cases := []struct {
		clause  Clause
		clause2 Clause
		want    bool
	}{
		{empty, one, true},
		{empty, two, true},
		{one, two, true},
		{one, three, true},
		{two, three, true},
		{one, four, false},
		{two, five, false},
	}

	for _, c := range cases {
		got := c.clause.Subsumes(c.clause2)
		if got != c.want {
			t.Errorf("%q.Subsumes(%q) == %t, want %t", c.clause, c.clause2, got, c.want)
		}
	}
}
