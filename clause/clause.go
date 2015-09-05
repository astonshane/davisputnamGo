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
	clause literalSlice
}

//Append adds a literal to the clause
func (c *Clause) Append(l literal.Literal) {
	//don't add the literal if it is already in the clause
	for _, lit := range c.clause {
		if l == lit {
			return
		}
	}
	c.clause = append(c.clause, l)
	sort.Sort(c.clause)
}

//Remove removes the given literal from the clause set
func (c *Clause) Remove(l literal.Literal) {
	newClause := literalSlice{}
	for _, lit := range c.clause {
		if l != lit {
			newClause = append(newClause, lit)
		}
	}
	c.clause = newClause
}

//RemoveIndex removes the literal at index i
func (c *Clause) RemoveIndex(i int) {
	if i < len(c.clause) {
		a := c.clause[:i]
		b := c.clause[i+1:]
		for _, lit := range b {
			a = append(a, lit)
		}
		c.clause = a
	}
}

//Equals compares the equality of two clauses
func Equals(a, b Clause) bool {
	if len(a.clause) != len(b.clause) {
		return false
	}
	for i := 0; i < len(a.clause); i++ {
		if a.clause[i] != b.clause[i] {
			return false
		}
	}
	return true
}

func (c Clause) String() string {
	ret := "{"
	for _, lit := range c.clause {
		ret = ret + fmt.Sprint(lit.String()) + ", "
	}
	ret = strings.Trim(ret, ", ")
	ret = ret + "}"
	return ret
}

//Len returns the length of the clause (ie. how many literals it contains)
func Len(c Clause) int {
	return len(c.clause)
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
