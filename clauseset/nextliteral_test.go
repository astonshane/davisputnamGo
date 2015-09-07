package clauseset

import (
	"davisputnam/clause"
	"davisputnam/literal"
	"testing"
)

func TestNextLiteral(t *testing.T) {
	empty, one, _, _ := clause.ConstructTestClauses()

	//case0 -- emptyClauseSet
	cs := ClauseSet{}
	got, err := cs.NextLiteral()
	want := literal.Literal{}
	if !literal.Equals(got, want) || err == nil {
		t.Errorf("%q.FirstElement() == %q, want %q (error: %q)", cs, got, want, err)
	}

	//case1 -- empty clause in emptyClauseSet
	cs = ClauseSet{}
	cs.Append(empty)
	got, err = cs.NextLiteral()
	want = literal.Literal{}
	if !literal.Equals(got, want) || err == nil {
		t.Errorf("%q.FirstElement() == %q, want %q (error: %q)", cs, got, want, err)
	}

	//case2 -- one clause in emptyClauseSet
	cs = ClauseSet{}
	cs.Append(one)
	got, err = cs.NextLiteral()
	want = literal.Literal{"A", false}
	if !literal.Equals(got, want) || err != nil {
		t.Errorf("%q.FirstElement() == %q, want %q (error: %q)", cs, got, want, err)
	}

}
