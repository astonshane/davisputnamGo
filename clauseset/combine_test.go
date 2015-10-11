package clauseset

import (
	"davisputnam/clause"
	"testing"
)

func TestCombine(t *testing.T) {
	empty, one, two, three := clause.ConstructTestClauses()
	cs0 := ClauseSet{}

	cs1 := ClauseSet{}
	cs1.Append(empty)

	cs2 := ClauseSet{}
	cs2.Append(one)

	cs3 := ClauseSet{}
	cs3.Append(two)
	cs3.Append(three)

	cases := []struct {
		a    ClauseSet
		b    ClauseSet
		want string
	}{
		{cs0, cs0, "{}"},
		{cs0, cs1, "{{}}"},
		{cs1, cs1, "{{}}"},
		{cs1, cs2, "{{}, {A}}"},
		{cs2, cs2, "{{A}}"},
		{cs2, cs3, "{{A}, {A, B}, {A, B, ~B}}"},
	}
	for _, clause := range cases {
		got := Combine(clause.a, clause.b).String()
		if got != clause.want {
			t.Errorf("Combine(%q, %q): %q != %q", clause.a, clause.b, got, clause.want)
		}
	}
}
