package literal

//Literal is a container class for a name and negation of a literal -- ex. A, ~A
type Literal struct {
	Name    string
	Negated bool
}

func (l Literal) String() string {
	if l.Negated {
		return "~" + l.Name
	}
	return l.Name
}

//Negation returns the negated literal
func (l *Literal) Negation() Literal {
	return Literal{Name: l.Name, Negated: !l.Negated}
}

//Equals compares two Literals by their output strings
func Equals(a, b Literal) bool {
	return a.String() == b.String()
}

//ConstructTestLiterals used by the tests
func ConstructTestLiterals() (Literal, Literal, Literal) {
	a := Literal{Name: "A", Negated: false}
	b := Literal{Name: "B", Negated: false}
	nb := Literal{Name: "B", Negated: true}
	return a, b, nb
}
