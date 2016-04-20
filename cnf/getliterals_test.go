package cnf

import (
	"testing"
)

func TestGetLiterals(t *testing.T) {

	cases := []struct {
		statement string
		want []string
	}{
		{"A", []string{"A"}},
		{"~A", []string{"A"}},
		{"~A<->B", []string{"A", "B"}},
		{"(~A^B)vQ", []string{"A", "B", "Q"}},
	}

	for _, c := range cases {
		got := getLiterals(c.statement)

		if len(got) != len(c.want) {
			t.Errorf("getLiterals(%s) == %t, want %t", c.statement, got, c.want)
		}
		for i, _ := range got {
			if got[i] != c.want[i] {
				t.Errorf("getLiterals(%s)[%d] == %t, want %t", c.statement, i, got[i], c.want[i])
			}
		}
	}
}
