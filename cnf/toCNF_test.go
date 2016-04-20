package cnf

import (
	"testing"
)

func TestToCNF(t *testing.T) {

	cases := []struct {
		statement string
		want int
	}{
		{"(AvB)vC", 8},
	}

	for _, c := range cases {
		toCNF(c.statement)
	}
}
