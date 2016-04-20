package cnf

import (
    "regexp"
    "fmt"
)

func getLiterals(str string) []string{
    literals := []string{}

    r, _ := regexp.Compile("([A-Z])")
    matches := r.FindAllString(str, -1)

    for _, l := range matches {
        literals = append(literals, l)
    }

    return literals
}

func getRows(literals []string) [][]string {
    cases := [][]string{}
    if len(literals) == 0 {
        return append(cases, []string{})
    }
    new_literal := literals[0]
    literals = literals[1:]

    new_cases := getRows(literals)
    for _, c := range new_cases {
        c1 := append(c, new_literal)
        c2 := append(c, fmt.Sprintf("~(%s)", new_literal))
        cases = append(cases, c1, c2)
    }

    return cases
}
