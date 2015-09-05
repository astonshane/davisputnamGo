package clause

import (
	"davisputnam/literal"
	"testing"
)

func TestRemoveIndex(t *testing.T) {
	a, b, nb := literal.ConstructTestLiterals()

	//case0 -- remove from empty clause
	clause := Clause{}
	before := clause.String()
	clause.Remove(a)
	got := clause.String()
	want := "{}"
	if got != want {
		t.Errorf("%q.Remove(%q) == %q, want %q", before, a, got, want)
	}

	//case1 -- remove only literal from the clause
	clause = Clause{}
	clause.Append(a)
	before = clause.String()
	clause.Remove(a)
	got = clause.String()
	want = "{}"
	if got != want {
		t.Errorf("%q.Remove(%q) == %q, want %q", before, a, got, want)
	}

	//case2 -- 2 different literals, remove first
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	before = clause.String()
	clause.Remove(a)
	got = clause.String()
	want = "{B}"
	if got != want {
		t.Errorf("%q.Remove(%q) == %q, want %q", before, a, got, want)
	}

	//case3 -- 2 different literals, remove second
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	before = clause.String()
	clause.Remove(b)
	got = clause.String()
	want = "{A}"
	if got != want {
		t.Errorf("%q.Remove(%q) == %q, want %q", before, b, got, want)
	}

	//case4 -- 3 different literals, remove first
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	clause.Append(nb)
	before = clause.String()
	clause.Remove(a)
	got = clause.String()
	want = "{B, ~B}"
	if got != want {
		t.Errorf("%q.Remove(%q) == %q, want %q", before, a, got, want)
	}

	//case5 -- 3 different literals, remove second
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	clause.Append(nb)
	before = clause.String()
	clause.Remove(b)
	got = clause.String()
	want = "{A, ~B}"
	if got != want {
		t.Errorf("%q.Remove(%q) == %q, want %q", before, b, got, want)
	}

	//case6 -- 3 different literals, remove third
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	clause.Append(nb)
	before = clause.String()
	clause.Remove(nb)
	got = clause.String()
	want = "{A, B}"
	if got != want {
		t.Errorf("%q.Remove(%q) == %q, want %q", before, nb, got, want)
	}
}
