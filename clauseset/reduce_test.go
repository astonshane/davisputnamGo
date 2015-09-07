package clauseset

import (
	"davisputnam/clause"
	"davisputnam/literal"
	"testing"
)

func TestReduce(t *testing.T) {

	a, b, nb := literal.ConstructTestLiterals()

	cases := []struct {
		lit  literal.Literal
		want string
	}{
		{a, "{{}}"},
		{b, "{{A}, {~A}}"},
		{a.Negation(), "{{}, {B}, {~B}}"},
		{nb, "{{A}, {~A}}"},
	}

	for _, c := range cases {
		_, one, two, _ := clause.ConstructTestClauses()

		nOne := clause.Clause{}
		nOne.Append(a.Negation())

		nTwo := clause.Clause{}
		nTwo.Append(a)
		nTwo.Append(nb)

		cs := ClauseSet{}
		cs.Append(one)
		cs.Append(nOne)
		cs.Append(two)
		cs.Append(nTwo)

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
