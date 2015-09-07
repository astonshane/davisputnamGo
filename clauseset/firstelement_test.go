package clauseset

import (
	"davisputnam/clause"
	//"davisputnam/literal"
	"testing"
)

func TestFirstElement(t *testing.T) {
	empty, one, _, _ := clause.ConstructTestClauses()

	//case0 -- emptyClauseSet
	cs := ClauseSet{}
	got, err := cs.FirstElement()
	want := clause.Clause{}

	if !clause.Equals(got, want) || err == nil {
		t.Errorf("%q.FirstElement() == %q, want %q (error: %q)", cs, got, want, err)
	}

	//case1 -- one thing
	cs = ClauseSet{}
	cs.Append(empty)
	got, err = cs.FirstElement()
	want = clause.Clause{}

	if !clause.Equals(got, want) || err != nil {
		t.Errorf("%q.FirstElement() == %q, want %q (error: %q)", cs, got, want, err)
	}

	//case2 -- two things
	cs = ClauseSet{}
	cs.Append(one)
	cs.Append(empty)
	got, err = cs.FirstElement()
	want = clause.Clause{}

	if !clause.Equals(got, want) || err != nil {
		t.Errorf("%q.FirstElement() == %q, want %q (error: %q)", cs, got, want, err)
	}
}
