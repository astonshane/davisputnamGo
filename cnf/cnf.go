package cnf

import (
    "regexp"
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
