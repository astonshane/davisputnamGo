package clauseset

import (
	"davisputnam/clause"
	"testing"
)

func TestAppend(t *testing.T) {
	empty, one, two, _ := clause.ConstructTestClauses()

	//case0 -- 1 clause
	cs := ClauseSet{}
	got := cs.String()
	cs.Append(empty)
	want := "{{}}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case1 -- 2 things in ClauseSet (added in sorted order)
	cs = ClauseSet{}
	cs.Append(empty)
	cs.Append(one)
	got = cs.String()
	want = "{{}, {A}}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case2 -- 2 things in ClauseSet (added in non-sorted order)
	cs = ClauseSet{}
	cs.Append(one)
	cs.Append(empty)
	got = cs.String()
	want = "{{}, {A}}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case3 -- 3 things in ClauseSet (added in sorted order)
	cs = ClauseSet{}
	cs.Append(empty)
	cs.Append(one)
	cs.Append(two)
	got = cs.String()
	want = "{{}, {A}, {A, B}}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case4 -- 3 things in ClauseSet (added in non-sorted order)
	cs = ClauseSet{}
	cs.Append(two)
	cs.Append(empty)
	cs.Append(one)
	got = cs.String()
	want = "{{}, {A}, {A, B}}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

}
