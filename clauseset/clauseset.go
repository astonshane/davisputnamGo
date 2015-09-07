package clauseset

import (
	"davisputnam/clause"
	"davisputnam/literal"
	"errors"
	"sort"
	"strings"
)

type clauseSlice []clause.Clause

//ClauseSet holds a slice of clauses
type ClauseSet struct {
	clauses clauseSlice
}

/*func (c ClauseSet) Copy() ClauseSet {
	new_clasuseset := c.clauses
	return new_clasuseset
}*/

//Append adds a clause to the ClauseSet
func (c *ClauseSet) Append(a clause.Clause) {
	for _, b := range c.clauses {
		if clause.Equals(a, b) {
			return
		}
	}
	c.clauses = append(c.clauses, a)
	sort.Sort(c.clauses)
}

//Reduce transforms S by removing any clauses that contain the literal L1
//   and removing L2 (i.e. ~L1) from any clauses that contain it
func (c ClauseSet) Reduce(l1 literal.Literal) ClauseSet {
	cs := c.Copy()
	l2 := l1.Negation()

	newClauseSet := ClauseSet{}

	for _, clause := range cs.clauses {
		i := clause.Contains(l1)
		if i != -1 {
			continue
		}

		j := clause.Contains(l2)
		if j != -1 {
			clause.RemoveIndex(j)
		}
		newClauseSet.Append(clause)

	}
	return newClauseSet
}

//FirstElement returns the first element in ClauseSet
func (c ClauseSet) FirstElement() (clause.Clause, error) {
	if c.Len() > 0 {
		return c.clauses[0], nil
	}
	//return empty clause + error
	return clause.Clause{}, errors.New("No first element in empty ClauseSet")
}

//NextLiteral returns the first literal in the ClauseSet that it finds
func (c ClauseSet) NextLiteral() (literal.Literal, error) {
	if c.Len() > 0 {
		for _, cla := range c.clauses {
			for _, lit := range cla.Clause {
				return lit.Copy(), nil
			}
		}
	}
	//return empty literal + literal
	return literal.Literal{}, errors.New("No literals in ClauseSet")

}

//Equals returns the equality of two ClauseSets
func Equals(a, b ClauseSet) bool {
	if a.Len() != b.Len() {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		if !clause.Equals(a.clauses[i], b.clauses[i]) {
			return false
		}
	}
	return true
}

//Len returns the length of the ClauseSet (ie. how many clauses it contains)
func Len(c ClauseSet) int {
	return len(c.clauses)
}

//Copy copies the contesnts of ClauseSet
func (c ClauseSet) Copy() ClauseSet {
	newClauseSet := ClauseSet{}

	for _, cl := range c.clauses {
		newClauseSet.Append(cl.Copy())
	}
	return newClauseSet
}

func (c ClauseSet) String() string {
	ret := "{"
	for _, clause := range c.clauses {
		ret = ret + clause.String() + ", "
	}
	ret = strings.Trim(ret, ", ")
	ret = ret + "}"
	return ret
}

//Len returns the length of the ClauseSet
func (c ClauseSet) Len() int {
	return len(c.clauses)
}

//functions needed to define the Sort interface for type clauseSlice([]clause.Clause)
func (c clauseSlice) Len() int {
	return len(c)
}
func (c clauseSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c clauseSlice) Less(i, j int) bool {
	//if they aren't the same length, put the shorter one first
	if clause.Len(c[i]) != clause.Len(c[j]) {
		return clause.Len(c[i]) < clause.Len(c[j])
	}

	//they must both be the same size now...
	for x := 0; x < clause.Len(c[i]); x++ {
		//literal x is not the same
		if c[i].Clause[x] != c[j].Clause[x] {
			//return the lesser one
			return literal.Less(c[i].Clause[x], c[j].Clause[x])
		}
	}

	return false

}
