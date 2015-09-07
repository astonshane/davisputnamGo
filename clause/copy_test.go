package clause

import (
	"davisputnam/literal"
	"testing"
)

func TestCopy(t *testing.T) {
	a, b, _ := literal.ConstructTestLiterals()

	//case0 -- empty Clause, don't change
	c := Clause{}
	copy := c.Copy()
	if !Equals(c, copy) {
		t.Errorf("%q.Copy() == %q, want %q", c, copy, c)
	}

	//case1 -- empty Clause, change copy
	c = Clause{}
	copy = c.Copy()
	copy.Append(a)
	if Equals(c, copy) {
		t.Errorf("copy (%q) == original (%q)", copy, c)
	}

	//case2 -- 1 thing in Clause, don't change
	c = Clause{}
	c.Append(a)
	copy = c.Copy()
	if !Equals(c, copy) {
		t.Errorf("%q.Copy() == %q, want %q", c, copy, c)
	}

	//case3 -- 1 thing in Clause, change copy
	c = Clause{}
	c.Append(a)
	copy = c.Copy()
	copy.Append(b)
	if Equals(c, copy) {
		t.Errorf("copy (%q) == original (%q)", copy, c)
	}

}
