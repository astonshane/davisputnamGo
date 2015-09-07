package clauseset

import (
	"davisputnam/clause"
	"testing"
)

func TestCopy(t *testing.T) {
	_, one, two, _ := clause.ConstructTestClauses()

	//case0 -- empty ClauseSet, don't change
	cs := ClauseSet{}
	copy := cs.Copy()

	if !Equals(cs, copy) {
		t.Errorf("original (%q) != copy (%q)", cs, copy)
	}

	//case1 -- empty ClauseSet, chagne copy
	cs = ClauseSet{}
	copy = cs.Copy()
	copy.Append(one)

	if Equals(cs, copy) {
		t.Errorf("original (%q) == Copy().Append(%q) (%q)", cs, one, copy)
	}

	//case2 -- 1 thing in ClauseSet, don't change
	cs = ClauseSet{}
	cs.Append(two)
	copy = cs.Copy()

	if !Equals(cs, copy) {
		t.Errorf("original (%q) != copy (%q)", cs, copy)
	}

	//case3 -- 1 thing in ClauseSet, chagne copy
	cs = ClauseSet{}
	cs.Append(two)
	copy = cs.Copy()
	copy.Append(one)

	if Equals(cs, copy) {
		t.Errorf("original (%q) == Copy().Append(%q, %q) (%q)", cs, one, two, copy)
	}

}
