package clause

import (
	"davisputnam/literal"
	"fmt"
	"sort"
	"strings"
)

type literalSlice []literal.Literal

//Clause is a slice of Literals -- ex. {A, ~B, C}, {~A}
type Clause struct {
	Clause literalSlice
}

//Append adds a literal to the clause
func (c *Clause) Append(l literal.Literal) {
	//don't add the literal if it is already in the clause
	for _, lit := range c.Clause {
		if l == lit {
			return
		}
	}
	c.Clause = append(c.Clause, l)
	sort.Sort(c.Clause)
}

//Remove removes the given literal from the clause set
func (c *Clause) Remove(l literal.Literal) {
	newClause := literalSlice{}
	for _, lit := range c.Clause {
		if l != lit {
			newClause = append(newClause, lit)
		}
	}
	c.Clause = newClause
}

//RemoveIndex removes the literal at index i
func (c *Clause) RemoveIndex(i int) {
	if i < len(c.Clause) {
		a := c.Clause[:i]
		b := c.Clause[i+1:]
		for _, lit := range b {
			a = append(a, lit)
		}
		c.Clause = a
	}
}

//Equals compares the equality of two clauses
func Equals(a, b Clause) bool {
	if len(a.Clause) != len(b.Clause) {
		return false
	}
	for i := 0; i < len(a.Clause); i++ {
		if a.Clause[i] != b.Clause[i] {
			return false
		}
	}
	return true
}

func (c Clause) String() string {
	ret := "{"
	for _, lit := range c.Clause {
		ret = ret + fmt.Sprint(lit.String()) + ", "
	}
	ret = strings.Trim(ret, ", ")
	ret = ret + "}"
	return ret
}

//Len returns the length of the clause (ie. how many literals it contains)
func Len(c Clause) int {
	return len(c.Clause)
}

//functions needed to define the Sort interface for type literalSlice ([]literal.Literal)
func (c literalSlice) Len() int {
	return len(c)
}
func (c literalSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c literalSlice) Less(i, j int) bool {
	if c[i].Name < c[j].Name {
		return true
	} else if c[i].Name > c[j].Name {
		return false
	} else {
		if c[i].Negated == c[j].Negated {
			return false
		} else if !c[i].Negated && c[j].Negated {
			return true
		}
		return false
	}
}

//ConstructTestClauses used by the tests
func ConstructTestClauses() (Clause, Clause, Clause, Clause) {
	a, b, nb := literal.ConstructTestLiterals()

	empty := Clause{}

	one := Clause{}
	one.Append(a)

	two := Clause{}
	two.Append(a)
	two.Append(b)

	three := Clause{}
	three.Append(a)
	three.Append(b)
	three.Append(nb)

	return empty, one, two, three

}
