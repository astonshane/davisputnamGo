package clause

import (
	"davisputnam/literal"
	"testing"
)

func TestPrint(t *testing.T) {
	a, b, nb := literal.ConstructTestLiterals()

	//case1 -- empty clause
	c1 := Clause{}
	got := c1.String()
	want := "{}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case2 -- 1 thing
	c2 := Clause{}
	c2.Append(a)
	got = c2.String()
	want = "{A}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case 3 -- 2 things
	c3 := Clause{}
	c3.Append(b)
	c3.Append(a)
	got = c3.String()
	want = "{A, B}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

	//case 4 -- 3 things
	c4 := Clause{}
	c4.Append(nb)
	c4.Append(b)
	c4.Append(a)
	got = c4.String()
	want = "{A, B, ~B}"
	if got != want {
		t.Errorf("String() == %q, want %q", got, want)
	}

}
