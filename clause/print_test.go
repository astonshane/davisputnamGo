package clause

import (
	"davisputnam/literal"
	"testing"
)

func ConstructLiterals() (literal.Literal, literal.Literal, literal.Literal) {
	a := literal.Literal{Name: "A", Negated: false}
	b := literal.Literal{Name: "B", Negated: true}
	c := literal.Literal{Name: "C", Negated: true}
	return a, b, c
}

func TestPrint(t *testing.T) {
	a, b, c := ConstructLiterals()

	//case1 -- empty clause set
	c1 := Clause{}
	got := c1.Print()
	want := "{}"
	if got != want {
		t.Errorf("Print() == %q, want %q", got, want)
	}

	//case2 -- 1 thing
	c2 := Clause{}
	c2.Append(a)
	got = c2.Print()
	want = "{A}"
	if got != want {
		t.Errorf("Print() == %q, want %q", got, want)
	}

	//case 3 -- 2 things
	c3 := Clause{}
	c3.Append(b)
	c3.Append(a)
	got = c3.Print()
	want = "{~B, A}"
	if got != want {
		t.Errorf("Print() == %q, want %q", got, want)
	}

	//case 4 -- 3 things
	c4 := Clause{}
	c4.Append(a)
	c4.Append(b)
	c4.Append(c)
	got = c4.Print()
	want = "{A, ~B, ~C}"
	if got != want {
		t.Errorf("Print() == %q, want %q", got, want)
	}

}
