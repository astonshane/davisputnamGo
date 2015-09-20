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

//Copy returns a copy of the literal
func (l Literal) Copy() Literal {
	return Literal{l.Name, l.Negated}
}

//Negation returns the negated literal
func (l *Literal) Negation() Literal {
	return Literal{Name: l.Name, Negated: !l.Negated}
}

//Equals compares two Literals by their output strings
func Equals(a, b Literal) bool {
	return a.String() == b.String()
}

//Less compares two literals, returns the "lower one" -- A < B;  A < ~A; A < ~B; A !< A
func Less(a, b Literal) bool {
	if a.Name == b.Name {
		//same name
		if a.Negated == b.Negated {
			//same negation too
			return false
		}
		//if a is not negated, it should be Less
		return !a.Negated
	}
	//different name, return a<b by name
	return a.Name < b.Name
}

//ConstructTestLiterals used by tests
func ConstructTestLiterals() (Literal, Literal, Literal) {
	a := Literal{Name: "A", Negated: false}
	b := Literal{Name: "B", Negated: false}
	nb := Literal{Name: "B", Negated: true}
	return a, b, nb
}

//ConstructMoreTestLiterals used by tests
func ConstructMoreTestLiterals() (Literal, Literal, Literal) {
	c := Literal{Name: "C", Negated: false}
	d := Literal{Name: "D", Negated: false}
	e := Literal{Name: "E", Negated: false}
	return c, d, e
}
