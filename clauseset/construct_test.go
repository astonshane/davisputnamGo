package clauseset

import (
	"davisputnam/connector"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {

	cases := []struct {
		plaintext string
		want      string
	}{
		{"A", "{{A}}"},
		{"B", "{{B}}"},
		{"~A", "{{~A}}"},
		{"~B", "{{~B}}"},
		{"A^B", "{{A}, {B}}"},
		{"A^~B", "{{A}, {~B}}"},
		{"~A^~B", "{{~A}, {~B}}"},
		{"A^~A^~B", "{{A}, {~A}, {~B}}"},
		{"AvB", "{{A, B}}"},
		{"Av~B", "{{A, ~B}}"},
		{"~Av~B", "{{~A, ~B}}"},
		{"Av~Av~B", "{{A, ~A, ~B}}"},
	}
	for _, c := range cases {
		parsed := connector.Parse(c.plaintext)
		fmt.Println(parsed)
		got := ConstructCS(parsed).String()

		want := c.want
		if got != want {
			t.Errorf("ConstructCS(Parse(%q)): %q != %q", c.plaintext, got, want)
		}
	}
}
