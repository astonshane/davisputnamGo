package connector

import "fmt"

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
