package clause

import (
	"davisputnam/literal"
	"testing"
)

func TestAppend(t *testing.T) {
	a, b, nb := literal.ConstructTestLiterals()

	//case0 -- 1 literal
	c0 := Clause{}
	c0.Append(a)
	got := c0.String()
	want := "{A}"
	if got != want {
		t.Errorf("Append(%q) == %q, want %q", a, got, want)
	}

	//case1 -- 2 of the same literal
	c1 := Clause{}
	c1.Append(a)
	c1.Append(a)
	got = c1.String()
	want = "{A}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", a, a, got, want)
	}

	//case2 -- 2 different literals (added in sorted order)
	c2 := Clause{}
	c2.Append(a)
	c2.Append(b)
	got = c2.String()
	want = "{A, B}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", a, b, got, want)
	}

	//case3 -- 2 different literals (added in non-sorted order)
	c3 := Clause{}
	c3.Append(a)
	c3.Append(b)
	got = c3.String()
	want = "{A, B}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", b, a, got, want)
	}

	//case4 -- 2 opposite literals (added in sorted order)
	c4 := Clause{}
	c4.Append(b)
	c4.Append(nb)
	got = c4.String()
	want = "{B, ~B}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", b, nb, got, want)
	}

	//case5 -- 2 opposite literals (added in non-sorted order)
	c5 := Clause{}
	c5.Append(nb)
	c5.Append(b)
	got = c5.String()
	want = "{B, ~B}"
	if got != want {
		t.Errorf("Append(%q, %q) == %q, want %q", nb, b, got, want)
	}

	//case6 -- 3  literals (added in sorted order)
	c6 := Clause{}
	c6.Append(a)
	c6.Append(b)
	c6.Append(nb)
	got = c6.String()
	want = "{A, B, ~B}"
	if got != want {
		t.Errorf("Append(%q, %q, %q) == %q, want %q", a, b, nb, got, want)
	}

	//case7 -- 3  literals (added in non-sorted order)
	c7 := Clause{}
	c7.Append(nb)
	c7.Append(b)
	c7.Append(a)
	got = c7.String()
	want = "{A, B, ~B}"
	if got != want {
		t.Errorf("Append(%q, %q, %q) == %q, want %q", nb, b, a, got, want)
	}

}
