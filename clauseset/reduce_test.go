package clauseset

import (
	"davisputnam/clause"
	"davisputnam/literal"
	"testing"
)

func TestReduce(t *testing.T) {

	a, b, _ := literal.ConstructTestLiterals()
	c := literal.Literal{Name: "C", Negated: false}
	one, two, three, four := clause.ConstructMoreTestClauses()
	clauseSet := ClauseSet{}
	clauseSet.Append(one)
	clauseSet.Append(two)
	clauseSet.Append(three)
	clauseSet.Append(four)

	cases := []struct {
		lit  literal.Literal
		want string
	}{
		{a, "{{B, C}}"},
		{b, "{{A}}"},
		{a.Negation(), "{{}, {B}}"},
		{c, "{{A}, {A, B}}"},
	}

	for _, c := range cases {
		cs := clauseSet.Copy()
		before := cs.String() //{{A}, {~A}, {A, B}, {A, ~B}}
		new := cs.Reduce(c.lit)
		post := cs.String()
		got := new.String()

		if got != c.want {
			t.Errorf("%q.Reduce(%q) == %q, want %q", before, c.lit, got, c.want)
		}
		if before != post {
			t.Errorf("before (%q) != post (%q); cs changed", before, post)
		}
	}
}
