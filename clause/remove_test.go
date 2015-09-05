package clause

import (
	"davisputnam/literal"
	"testing"
)

func constructLiterals3() (literal.Literal, literal.Literal, literal.Literal) {
	a := literal.Literal{Name: "A", Negated: false}
	b := literal.Literal{Name: "B", Negated: false}
	nb := literal.Literal{Name: "B", Negated: true}
	return a, b, nb
}

func TestRemove(t *testing.T) {
	a, b, nb := constructLiterals3()

	//case0 -- remove from empty clause
	clause := Clause{}
	before := clause.String()
	clause.RemoveIndex(1)
	got := clause.String()
	want := "{}"
	if got != want {
		t.Errorf("%q.RemoveIndex(1) == %q, want %q", before, got, want)
	}

	//case1 -- remove only literal from the clause
	clause = Clause{}
	clause.Append(a)
	before = clause.String()
	clause.RemoveIndex(0)
	got = clause.String()
	want = "{}"
	if got != want {
		t.Errorf("%q.RemoveIndex(0) == %q, want %q", before, got, want)
	}

	//case2 -- 2 different literals, remove first
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	before = clause.String()
	clause.RemoveIndex(0)
	got = clause.String()
	want = "{B}"
	if got != want {
		t.Errorf("%q.RemoveIndex(0) == %q, want %q", before, got, want)
	}

	//case3 -- 2 different literals, remove second
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	before = clause.String()
	clause.RemoveIndex(1)
	got = clause.String()
	want = "{A}"
	if got != want {
		t.Errorf("%q.RemoveIndex(1) == %q, want %q", before, got, want)
	}

	//case4 -- 3 different literals, remove first
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	clause.Append(nb)
	before = clause.String()
	clause.RemoveIndex(0)
	got = clause.String()
	want = "{B, ~B}"
	if got != want {
		t.Errorf("%q.RemoveIndex(0) == %q, want %q", before, got, want)
	}

	//case5 -- 3 different literals, remove second
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	clause.Append(nb)
	before = clause.String()
	clause.RemoveIndex(1)
	got = clause.String()
	want = "{A, ~B}"
	if got != want {
		t.Errorf("%q.RemoveIndex(1) == %q, want %q", before, got, want)
	}

	//case6 -- 3 different literals, remove third
	clause = Clause{}
	clause.Append(a)
	clause.Append(b)
	clause.Append(nb)
	before = clause.String()
	clause.RemoveIndex(2)
	got = clause.String()
	want = "{A, B}"
	if got != want {
		t.Errorf("%q.RemoveIndex(2) == %q, want %q", before, got, want)
	}
}
