package main

import (
	"davisputnam/clause"
	"davisputnam/literal"
	"fmt"
)

func main() {
	l := literal.Literal{Name: "A", Negated: false}
	m := literal.Literal{Name: "C", Negated: true}

	c := clause.Clause{}
	c.Append(l)
	c.Append(m)

	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(c)
}
