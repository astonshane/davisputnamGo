package clauseset

import (
	"davisputnam/clause"
	"davisputnam/literal"
	"sort"
	"strings"
)

type clauseSlice []clause.Clause

//ClauseSet holds a slice of clauses
type ClauseSet struct {
	clauses clauseSlice
}

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

func (c ClauseSet) String() string {
	ret := "{"
	for _, clause := range c.clauses {
		ret = ret + clause.String() + ", "
	}
	ret = strings.Trim(ret, ", ")
	ret = ret + "}"
	return ret
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
