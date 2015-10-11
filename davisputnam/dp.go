package main

import (
	"davisputnam/clause"
	"davisputnam/clauseset"
	"davisputnam/literal"
	"fmt"
	"log"
	"strings"
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
	fmt.Printf("\n%sSatisfiable(%s)\n", strings.Repeat(" ", CS.Indent), CS)
	CS.Indent += 2

	//if CS = {} : return true
	if CS.Len() == 0 {
		fmt.Printf("%sEmpty ClauseSet, returning true\n", strings.Repeat(" ", CS.Indent))
		return true
	}

	//if {} in CS : return false
	firstElement, err := CS.FirstElement()
	if err != nil {
		log.Fatal(err)
	}
	if clause.Len(firstElement) == 0 {
		fmt.Printf("%s{} found in ClauseSet, returning false\n", strings.Repeat(" ", CS.Indent))
		return false
	}

	//select L in lit(CS): return Satisfiable(CS_L) || Satisfiable(CS_L')
	nextLiteral, err2 := CS.NextLiteral()
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("%sSpliting on %s\n", strings.Repeat(" ", CS.Indent), nextLiteral)

	CSL := CS.Reduce(nextLiteral)
	CSL.Indent = CS.Indent
	CSR := CS.Reduce(nextLiteral.Negation())
	CSR.Indent = CS.Indent

	return Satisfiable(CSL) || Satisfiable(CSR)
}

//FindValidity finds the validity of the argument (given as a ClauseSet)
func FindValidity(CS clauseset.ClauseSet) {
	fmt.Printf("Starting ClauseSet: %s\n", CS)

	CS.Indent = 0
	sat := Satisfiable(CS)

	fmt.Print("Conclusion: ")
	if sat {
		fmt.Printf("Satisfiable(%s) == %t; INVALID\n", CS, sat)
	} else {
		fmt.Printf("Satisfiable(%s) == %t; VALID\n", CS, sat)
	}
}

func main() {
	//from ValidTest:
	a := literal.Literal{"A", false}
	na := a.Negation()
	n := literal.Literal{"N", false}
	nn := n.Negation()
	q := literal.Literal{"Q", false}
	nq := q.Negation()

	one := clause.Clause{}
	one.Append(a)

	two := clause.Clause{}
	two.Append(nq)

	three := clause.Clause{}
	three.Append(nn)

	four := clause.Clause{}
	four.Append(na)
	four.Append(n)
	four.Append(q)

	CS := clauseset.ClauseSet{}
	CS.Append(one)
	CS.Append(two)
	CS.Append(three)
	CS.Append(four)

	FindValidity(CS)

}
