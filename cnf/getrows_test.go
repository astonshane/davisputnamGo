package cnf

import (
	"testing"
)

func TestGetRows(t *testing.T) {

	cases := []struct {
		statement string
		want int
	}{
		{"A", 2},
		{"AvB", 4},
		{"(AvB)^C", 8},
	}

	for _, c := range cases {
		got := getRows(getLiterals(c.statement))
		//fmt.Println(got)
		if len(got) != c.want {
			t.Errorf("%len(getRows(getLiterals(%s))) == %i, want %i", c.statement, got, c.want)
		}
	}
}
