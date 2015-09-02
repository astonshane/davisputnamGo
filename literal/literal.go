package literal

//Literal is a literal
type Literal struct {
	Name    string
	Negated bool
}

//Print returns a string that represents the literal with ~ for a negation
func (l *Literal) Print() string {
	if l.Negated {
		return "~" + l.Name
	}
	return l.Name
}

//Negation returns the negated literal
func (l *Literal) Negation() Literal {
	return Literal{Name: l.Name, Negated: !l.Negated}
}
