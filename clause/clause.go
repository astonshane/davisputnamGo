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

//Print returns the string representation of the clause
func (c *Clause) Print() string {
	ret := "{"
	for _, lit := range c.clause {
		ret = ret + fmt.Sprint(lit.Print()) + ", "
	}
	ret = strings.Trim(ret, ", ")
	ret = ret + "}"
	return ret
}

//Append adds a literal to the clause
func (c *Clause) Append(l literal.Literal) {
	c.clause = append(c.clause, l)
}

func main() {
	l := literal.Literal{Name: "A", Negated: false}
	m := literal.Literal{Name: "B", Negated: true}

	c := Clause{}

	fmt.Println(c.Print())

	c.Append(l)
	c.Append(m)

	fmt.Println(c.Print())

}
