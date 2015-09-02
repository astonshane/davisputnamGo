package main

import (
	"davisputnam/literal"
	"fmt"
)

func main() {
	l := literal.Literal{Name: "A", Negated: false}
	fmt.Println(l.Print())
}
