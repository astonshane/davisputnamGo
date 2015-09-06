package clauseset

import (
	"davisputnam/clause"
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
	return clause.Len(c[i]) < clause.Len(c[j])
}
