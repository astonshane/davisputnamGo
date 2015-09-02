package clause

import (
	"davisputnam/literal"
	"fmt"
	"sort"
	"strings"
)

type literals []literal.Literal

//Clause is a set of Literals -- ex. {A, ~B, C}, {~A}
type Clause struct {
	clause literals
}

//Append adds a literal to the clause
func (c *Clause) Append(l literal.Literal) {
	//don't add the literal if it is already in the clause
	for _, lit := range c.clause {
		if literal.Equals(l, lit) {
			return
		}
	}
	c.clause = append(c.clause, l)
	sort.Sort(c.clause)
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

func (c literals) Len() int {
	return len(c)
}
func (c literals) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c literals) Less(i, j int) bool {
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
	//return c.clause[i] < c.clause[j]
}
