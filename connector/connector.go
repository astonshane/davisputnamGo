package connector

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

//isLiteral returns true if the connector is a literal or a negation of a literal
func (c Connector) isLiteral() bool {
	if c.Type == "Literal" {
		return true
	} else if c.Type == "Neg" {
		return len(c.Children) == 1 && c.Children[0].Type == "Literal"
	}
	return false

}

//isOr returns true if the connector is an Or of just literals or negations of literals
// OR if the connector is just a literal (ie. something ORd with nothing)
func (c Connector) isSimpleOr() bool {
	if c.Type == "Or" {
		//make sure each child is also a literal
		for _, child := range c.Children {
			if !child.isLiteral() {
				return false
			}
		}
		return true
	}
	//its not an OR, is it just a literal?
	return c.isLiteral()
}

//isCNF returns true if the conector is in CNF form
func (c Connector) isCNF() bool {
	if c.Type == "And" {
		for _, child := range c.Children {
			if !child.isSimpleOr() {
				return false
			}
		}
		return true
	}
	return c.isSimpleOr()
}

//ToCNF takes a connector that is not in CNF and returns it in CNF form
func (c Connector) ToCNF() Connector {
  c = c.RemoveImp()
  c = c.RemoveEquiv()
	return c
}

//Negate takes a connector and adds a negation to the front
//handles double negations, demorgans, etc.
func (c Connector) Negate() Connector {
	newc := Connector{Type: "Neg"}
	newc.Children = append(newc.Children, c)
	return newc
}

//RemoveImp removes an implication from a connector
//and returns the equivillent connector
func (c Connector) RemoveImp() Connector {
	if c.Type != "Imp"{
		return c
	}
	//A->B == ~AvB
	a := c.Children[0]
	na := a.Negate().PropagateNegations()
	b := c.Children[1]
	newc := Connector{Type: "Or"}
	newc.Children = append(newc.Children, na)
	newc.Children = append(newc.Children, b)

  return newc
}

//RemoveEquiv removes an equivillence from a connector
//and returns the equivillent connector
func (c Connector) RemoveEquiv() Connector {
	if c.Type != "Equiv"{
		return c
	}
	//A<->B == (A->B) ^ (B->A)
	a := c.Children[0]
	b := c.Children[1]

	new1 := Connector{Type: "Imp"}
	new1.Children = append(new1.Children, a)
	new1.Children = append(new1.Children, b)
	new1 = new1.RemoveImp()

	new2 := Connector{Type: "Imp"}
	new2.Children = append(new2.Children, b)
	new2.Children = append(new2.Children, a)
	new2 = new2.RemoveImp()

	newc := Connector{Type: "And"}
	newc.Children = append(newc.Children, new1)
	newc.Children = append(newc.Children, new2)
  return newc
}

//PropagateNegations takes a Connector and pushes all negations as far in as possible
func (c Connector) PropagateNegations() Connector {
	if c.Type == "Neg" {
		child := c.Children[0]

		if child.Type == "Neg" { //double negation
			return child.Children[0]
		} else if child.Type == "Literal" { //can't go any farther than negating a Literal
			return c
		}

		newC := Connector{}
		if child.Type == "And" {
			newC.Type = "Or"
		} else if child.Type == "Or" {
			newC.Type = "And"
		}

		newChildren := []Connector{}
		for _, superchild := range child.Children {
			newChildren = append(newChildren, superchild.Negate())
		}
		newC.Children = newChildren

		return newC.PropagateNegations()
	}

	newChildren := []Connector{}
	for _, child := range c.Children {
		newChildren = append(newChildren, child.PropagateNegations())
	}
	return Connector{Type: c.Type, Children: newChildren}

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

	connectors := []struct {
		short string
		full  string
	}{
		{"^", "And"},
		{"v", "Or"},
		{"<->", "Equiv"},
		{"->", "Imp"},
	}

	//simple cases: no parens...
	if !strings.Contains(plaintext, "(") {
		for _, connector := range connectors {
			if strings.Contains(plaintext, connector.short) {
				splitPlain := strings.Split(plaintext, connector.short)
				children := []Connector{}
				for _, child := range splitPlain {
					children = append(children, Parse(child))
				}
				return Connector{Type: connector.full, Children: children}
			}
		}
	}

	return Connector{}
}
