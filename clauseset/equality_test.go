package clauseset

import (
	"davisputnam/clause"
	"testing"
)

func TestEquality(t *testing.T) {
	_, one, two, _ := clause.ConstructTestClauses()

	//case0 -- empty ClauseSets
	a := ClauseSet{}
	b := ClauseSet{}

	if !Equals(a, b) {
		t.Errorf("%q != %q", a, b)
	}

	//case1 -- one thing
	a = ClauseSet{}
	a.Append(one)
	b = ClauseSet{}
	b.Append(one)

	if !Equals(a, b) {
		t.Errorf("%q != %q", a, b)
	}

	//case2 -- two things
	a = ClauseSet{}
	a.Append(one)
	a.Append(two)
	b = ClauseSet{}
	b.Append(one)
	b.Append(two)

	if !Equals(a, b) {
		t.Errorf("%q != %q", a, b)
	}

}
