package clauseset

import (
	"davisputnam/clause"
	"testing"
)

func TestPrint(t *testing.T) {
	_, one, two, three := clause.ConstructTestClauses()

	//case0 -- empty ClauseSet
	cs := ClauseSet{}
	got := cs.String()
	want := "{}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case1 -- 1 thing in ClauseSet
	cs = ClauseSet{}
	cs.Append(one)
	got = cs.String()
	want = "{{A}}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case2-- 2 things in ClauseSet
	cs = ClauseSet{}
	cs.Append(one)
	cs.Append(two)
	got = cs.String()
	want = "{{A}, {A, B}}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case3-- 3 things in ClauseSet
	cs = ClauseSet{}
	cs.Append(one)
	cs.Append(two)
	cs.Append(three)
	got = cs.String()
	want = "{{A}, {A, B}, {A, B, ~B}}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

}
