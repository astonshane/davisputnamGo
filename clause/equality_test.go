package clause

import (
	"davisputnam/literal"
	"testing"
)

func TestEquality(t *testing.T) {
	a, b, nb := literal.ConstructTestLiterals()

	//case0 -- empty clauses
	c0 := Clause{}
	c1 := Clause{}

	got := Equals(c0, c1)
	want := true
	if got != want {
		t.Errorf("Equals(%q, %q) == %t, want %t", c0, c1, got, want)
	}

	//case1 -- one thing in each (true)
	c0 = Clause{}
	c0.Append(a)
	c1 = Clause{}
	c1.Append(a)

	got = Equals(c0, c1)
	want = true
	if got != want {
		t.Errorf("Equals(%q, %q) == %t, want %t", c0, c1, got, want)
	}

	//case2 -- one thing in each (false)
	c0 = Clause{}
	c0.Append(a)
	c1 = Clause{}
	c1.Append(b)

	got = Equals(c0, c1)
	want = false
	if got != want {
		t.Errorf("Equals(%q, %q) == %t, want %t", c0, c1, got, want)
	}

	//case3 -- two things in each (true)
	c0 = Clause{}
	c0.Append(a)
	c0.Append(b)
	c1 = Clause{}
	c1.Append(a)
	c1.Append(b)

	got = Equals(c0, c1)
	want = true
	if got != want {
		t.Errorf("Equals(%q, %q) == %t, want %t", c0, c1, got, want)
	}

	//case4 -- two things in each (false)
	c0 = Clause{}
	c0.Append(a)
	c0.Append(b)
	c1 = Clause{}
	c1.Append(a)
	c1.Append(nb)

	got = Equals(c0, c1)
	want = false
	if got != want {
		t.Errorf("Equals(%q, %q) == %t, want %t", c0, c1, got, want)
	}

}
