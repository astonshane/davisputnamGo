package main

import (
	"davisputnam/clause"
	"davisputnam/clauseset"
	"davisputnam/literal"
	"fmt"
	"strings"
)

//ConvertLiteral takes a string and returns its literal
func ConvertLiteral(plaintext string) literal.Literal {
	//special case: one literal, A
	if len(plaintext) == 1 {
		return literal.Literal{Name: plaintext}
		//special case: one literal, negated, ~A
	} else if len(plaintext) == 2 {
		return literal.Literal{Name: strings.Trim(plaintext, "~"), Negated: true}
	}
	return literal.Literal{}
}

//Convert takes a plaintext input and transforms it to a CNF ClauseSet
func Convert(plaintext string) clauseset.ClauseSet {

	newClauseSet := clauseset.ClauseSet{}

	if len(plaintext) == 1 || len(plaintext) == 2 {
		lit := ConvertLiteral(plaintext)
		newClause := clause.Clause{}
		newClause.Append(lit)
		newClauseSet.Append(newClause)
		return newClauseSet
	}

	//there's no parenthesies, so it should be easier to parse
	if !strings.Contains(plaintext, "(") {
		if strings.Contains(plaintext, "v") { //or statement
			split := strings.Split(plaintext, "v")
			newClause := clause.Clause{}
			for _, l := range split {
				newClause.Append(ConvertLiteral(strings.Trim(l, " ")))
			}
			newClauseSet.Append(newClause)
			return newClauseSet
		} else if strings.Contains(plaintext, "^") { //and statement
			split := strings.Split(plaintext, "^")
			for _, l := range split {
				newClause := clause.Clause{}
				newClause.Append(ConvertLiteral(strings.Trim(l, " ")))
				newClauseSet.Append(newClause)
			}
			return newClauseSet
		}
	} else {
		fmt.Println("parens")
	}
	return newClauseSet
}

func main() {
	cases := []struct {
		plaintext string
		want      string
	}{
		{"A", "{{A}}"},
		{"~A", "{{~A}}"},
		{"A v B", "{{A, B}}"},
		{"A v B v C", "{{A, B, C}}"},
		{"~A v B", "{{~A, B}}"},
		{"A v B v ~C", "{{A, B, ~C}}"},
		{"A ^ B", "{{A}, {B}}"},
		{"A ^ B ^ C", "{{A}, {B}, {C}}"},
		{"~A ^ B", "{{~A}, {B}}"},
		{"A ^ B ^ ~C", "{{A}, {B}, {~C}}"},
	}

	for _, c := range cases {
		if Convert(c.plaintext).String() != c.want {
			fmt.Printf("Convert(%q) == %q, not %q", c.plaintext, Convert(c.plaintext), c.want)
		}
	}
}
