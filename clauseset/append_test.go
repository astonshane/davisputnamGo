package clauseset

import (
	"davisputnam/clause"
	"davisputnam/literal"
	"testing"
)

func TestAppend(t *testing.T) {
	empty, one, two, _ := clause.ConstructTestClauses()

	//case0 -- 1 clause
	cs := ClauseSet{}
	cs.Append(empty)
	got := cs.String()
	want := "{{}}"
	if got != want {
		t.Errorf("Append(%q) == %q, want %q", empty, got, want)
	}

	//case1 -- 2 things in ClauseSet (added in sorted order)
	cs = ClauseSet{}
	cs.Append(empty)
	cs.Append(one)
	got = cs.String()
	want = "{{}, {A}}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", empty, one, got, want)
	}

	//case2 -- 2 things in ClauseSet (added in non-sorted order)
	cs = ClauseSet{}
	cs.Append(one)
	cs.Append(empty)
	got = cs.String()
	want = "{{}, {A}}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", one, empty, got, want)
	}

	//case3 -- 3 things in ClauseSet (added in sorted order)
	cs = ClauseSet{}
	cs.Append(empty)
	cs.Append(one)
	cs.Append(two)
	got = cs.String()
	want = "{{}, {A}, {A, B}}"
	if got != want {
		t.Errorf("Append(%q, %q, %q) == %q, want %q", empty, one, two, got, want)
	}

	//case4 -- 3 things in ClauseSet (added in non-sorted order)
	cs = ClauseSet{}
	cs.Append(two)
	cs.Append(empty)
	cs.Append(one)
	got = cs.String()
	want = "{{}, {A}, {A, B}}"
	if got != want {
		t.Errorf("Append(%q, %q, %q) == %q, want %q", two, empty, one, got, want)
	}

	a, _, nb := literal.ConstructTestLiterals()
	nOne := clause.Clause{}
	nOne.Append(a.Negation())

	nTwo := clause.Clause{}
	nTwo.Append(a)
	nTwo.Append(nb)

	//case5 -- 2 things in ClauseSet -- same length, but one less than other (added in sorted order)
	cs = ClauseSet{}
	cs.Append(one)
	cs.Append(nOne)
	got = cs.String()
	want = "{{A}, {~A}}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", one, nOne, got, want)
	}

	//case6 -- 2 things in ClauseSet -- same length, but one less than other (added in non-sorted order)
	cs = ClauseSet{}
	cs.Append(nOne)
	cs.Append(one)
	got = cs.String()
	want = "{{A}, {~A}}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", nOne, one, got, want)
	}

	//case7 -- 2 things in ClauseSet -- same length, but one less than other (added in sorted order)
	cs = ClauseSet{}
	cs.Append(two)
	cs.Append(nTwo)
	got = cs.String()
	want = "{{A, B}, {A, ~B}}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", two, nTwo, got, want)
	}

	//case7 -- 2 things in ClauseSet -- same length, but one less than other (added in non-sorted order)
	cs = ClauseSet{}
	cs.Append(nTwo)
	cs.Append(two)
	got = cs.String()
	want = "{{A, B}, {A, ~B}}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", nTwo, two, got, want)
	}

}
