package main

import (
	"davisputnam/clause"
	"davisputnam/clauseset"
	"davisputnam/literal"
	"fmt"
	"log"
)

/*
def Satisfiable(CS):
    if CS = {}:
        return true
    if {} in CS:
        return false
    if {L} in CS:
        return Satisfiable(CS_L)
    select L in lit(CS)
        return Satisfiable(CS_L) | Satisfiable(CS_L')
*/

//Satisfiable implements above function
func Satisfiable(CS clauseset.ClauseSet) bool {
	//if CS = {} : return true
	if CS.Len() == 0 {
		return true
	}

	//if {} in CS : return false
	firstElement, err := CS.FirstElement()
	if err != nil {
		log.Fatal(err)
	}
	if clause.Len(firstElement) == 0 {
		return false
	}

	//select L in lit(CS): return Satisfiable(CS_L) || Satisfiable(CS_L')
	nextLiteral, err2 := CS.NextLiteral()
	if err2 != nil {
		log.Fatal(err2)
	}
	CSL := CS.Reduce(nextLiteral)
	CSR := CS.Reduce(nextLiteral.Negation())

	return Satisfiable(CSL) || Satisfiable(CSR)
}

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
