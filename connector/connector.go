package main

import (
	"fmt"
	"strings"
)

//Connector is a class to form a parseable structure from
type Connector struct {
	Type     string
	Literal  string
	Children []Connector
}

//String returns the string representation of the Connector, recursively
func (c Connector) String() string {
	if c.Type == "Literal" {
		return fmt.Sprintf("{%s: %s}", c.Type, c.Literal)
	} else if c.Type == "Neg" && len(c.Children) == 1 {
		return fmt.Sprintf("{%s: %s}", c.Type, c.Children[0])
	}
	return fmt.Sprintf("{%s: %s}", c.Type, c.Children)
}

//Parse parses a plaintext line into a Connector sequence
func Parse(plaintext string) Connector {
	plaintext = strings.Replace(plaintext, " ", "", -1)

	//special cases for dealing with literals (or negations of literals)
	if len(plaintext) == 1 {
		return Connector{Type: "Literal", Literal: plaintext}
	} else if len(plaintext) == 2 {
		plaintext = strings.Trim(plaintext, "~")
		return Connector{Type: "Neg", Children: []Connector{Connector{Type: "Literal", Literal: plaintext}}}
	}

	connectors := make(map[string]string)
	connectors["^"] = "And"
	connectors["v"] = "Or"
	connectors["<->"] = "Equiv"
	connectors["->"] = "Imp"

	//simple cases: no parens...
	if !strings.Contains(plaintext, "(") {

		for key, op := range connectors {
			if strings.Contains(plaintext, key) {
				splitPlain := strings.Split(plaintext, key)
				children := []Connector{}
				for _, child := range splitPlain {
					children = append(children, Parse(child))
				}
				return Connector{Type: op, Children: children}
			}
		}
	}

	return Connector{}
}

func main() {
	cases := []string{"A", "B", "~A", "A^B", "~A^B", "AvB", "Av~B", "A->B", "A<->B"}
	for _, c := range cases {
		fmt.Println(Parse(c))
	}
}
