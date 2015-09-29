package connector

import "testing"

func TestString(t *testing.T) {
	a := Connector{Type: "Literal", Literal: "A"}
	na := Connector{Type: "Neg", Children: []Connector{a}}

	b := Connector{Type: "Literal", Literal: "B"}
	nb := Connector{Type: "Neg", Children: []Connector{b}}

	aANDb := Connector{Type: "And", Children: []Connector{a, b}}

	cases := []struct {
		con  Connector
		want string
	}{
		{a, "{Literal: A}"},
		{b, "{Literal: B}"},
		{na, "{Neg: {Literal: A}}"},
		{nb, "{Neg: {Literal: B}}"},
		{Connector{Type: "And", Children: []Connector{a, b}}, "{And: [{Literal: A} {Literal: B}]}"},
		{Connector{Type: "And", Children: []Connector{a, nb}}, "{And: [{Literal: A} {Neg: {Literal: B}}]}"},
		{Connector{Type: "And", Children: []Connector{na, nb}}, "{And: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{Connector{Type: "And", Children: []Connector{a, na, nb}}, "{And: [{Literal: A} {Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{Connector{Type: "Or", Children: []Connector{a, b}}, "{Or: [{Literal: A} {Literal: B}]}"},
		{Connector{Type: "Or", Children: []Connector{a, nb}}, "{Or: [{Literal: A} {Neg: {Literal: B}}]}"},
		{Connector{Type: "Or", Children: []Connector{na, nb}}, "{Or: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{Connector{Type: "Or", Children: []Connector{a, na, nb}}, "{Or: [{Literal: A} {Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{Connector{Type: "Imp", Children: []Connector{a, b}}, "{Imp: [{Literal: A} {Literal: B}]}"},
		{Connector{Type: "Imp", Children: []Connector{a, nb}}, "{Imp: [{Literal: A} {Neg: {Literal: B}}]}"},
		{Connector{Type: "Imp", Children: []Connector{na, nb}}, "{Imp: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{Connector{Type: "Equiv", Children: []Connector{a, b}}, "{Equiv: [{Literal: A} {Literal: B}]}"},
		{Connector{Type: "Equiv", Children: []Connector{a, nb}}, "{Equiv: [{Literal: A} {Neg: {Literal: B}}]}"},
		{Connector{Type: "Equiv", Children: []Connector{na, nb}}, "{Equiv: [{Neg: {Literal: A}} {Neg: {Literal: B}}]}"},
		{Connector{Type: "And", Children: []Connector{a, aANDb}}, "{And: [{Literal: A} {And: [{Literal: A} {Literal: B}]}]}"},
	}
	for _, c := range cases {
		got := c.con.String()
		want := c.want
		if got != want {
			t.Errorf("String(): %q != %q", got, want)
		}
	}

}
