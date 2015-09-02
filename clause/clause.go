package clause

import (
	"davisputnam/literal"
	"fmt"
	"strings"
)

//Clause is a set of Literals -- ex. {A, ~B, C}, {~A}
type Clause struct {
	clause []literal.Literal
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

//Append adds a literal to the clause
func (c *Clause) Append(l literal.Literal) {
	c.clause = append(c.clause, l)
}
