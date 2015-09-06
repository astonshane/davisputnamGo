package main

import (
	"davisputnam/clause"
	"davisputnam/literal"
	"fmt"
)

/*
def Satisfiable(s):
    if S = {}:
        return true
    if S = {{}}:
        return false
    if {} in S:
        return false
    if {L} in S:
        return Satisfiable(S_L)
    select L in lit(s)
        return Satisfiable(S_L) | Satisfiable(S_L')
*/

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
