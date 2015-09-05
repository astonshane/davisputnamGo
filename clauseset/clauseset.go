package clauseset

import (
	"davisputnam/clause"
	"sort"
)

type clauseSlice []clause.Clause

//ClauseSet holds a slice of clauses
type ClauseSet struct {
	clauses clauseSlice
}

//Append adds a clause to the ClauseSet
func (c *ClauseSet) Append(cl clause.Clause) {
	//modify to check for duplicates (ala clause.Append())
	//add tests
	c.clauses = append(c.clauses, cl)
	sort.Sort(c.clauses)
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
