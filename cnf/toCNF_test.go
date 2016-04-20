package cnf

import (
	"testing"
)

func TestToCNF(t *testing.T) {

	cases := []struct {
		statement string
		want int
	}{
		{"A<->B", 8},
	}

	for _, c := range cases {
		ToCNF(c.statement)
	}
}
