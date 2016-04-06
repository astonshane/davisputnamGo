package main

import (
	"davisputnam/clause"
	"davisputnam/clauseset"
	"davisputnam/literal"
	"testing"
)

///////////////////////////////////////
/*
A -> (N v Q)
~(N v ~A)
A -> Q (conclusion)
*/
func SatValid(t *testing.T) {
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

	sat, _ := Satisfiable(CS)
	if sat {
		t.Errorf("ClauseSet %q was Satisfiable, so it is an invalid argument (it should be valid)", CS)
	}

}

///////////////////////////////////////
/*
A -> B
~A
~B (conclusion)
*/
func SatInvalid(t *testing.T) {
	na := literal.Literal{"A", true}
	b := literal.Literal{"B", false}

	one := clause.Clause{}
	one.Append(na)

	two := clause.Clause{}
	two.Append(b)

	three := clause.Clause{}
	three.Append(na)
	three.Append(b)

	CS := clauseset.ClauseSet{}
	CS.Append(one)
	CS.Append(two)
	CS.Append(three)
	CS.Indent = 0

	sat, _ := Satisfiable(CS)
	if !sat {
		t.Errorf("ClauseSet %q was not Satisfiable, so it is a valid argument (it should be invalid)", CS)
	}
}

func TestSatisfiable(t *testing.T) {
	SatValid(t)
	SatInvalid(t)
}
