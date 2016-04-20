package cnf

import (
    "regexp"
    "fmt"
    "sort"
    "davisputnam/literal"
    "davisputnam/clause"
    "davisputnam/clauseset"
)

type ByLength []string
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) > len(s[j])
}


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

func toCNF(statement string) {

    rows := getRows(getLiterals(statement))
    fmt.Println(statement)
    new_clauseset := clauseset.ClauseSet{}
    for _, row := range rows {
        sort.Sort(ByLength(row))
        //fmt.Println(row)

        statement_cpy := statement[:1]+statement[1:]
        for _, l := range row {
            // @ == TRUE        "!" == FALSE
            replacement_string := "N/A"
            if len(l) == 1{
                replacement_string = "@"
            }else{
                replacement_string = "!"
                l = l[2:3]
            }
            // do the straight replacement
            r := regexp.MustCompile(fmt.Sprintf("%s", l))
            statement_cpy = r.ReplaceAllString(statement_cpy, replacement_string)

        }
        for statement_cpy != "TRUE" && statement_cpy != "FALSE"{
            r := regexp.MustCompile("(~@|~\\(@\\))")
            statement_cpy = r.ReplaceAllString(statement_cpy, "FALSE")

            r = regexp.MustCompile("(~!|~\\(!\\))")
            statement_cpy = r.ReplaceAllString(statement_cpy, "TRUE")

            r = regexp.MustCompile("\\(@\\)")
            statement_cpy = r.ReplaceAllString(statement_cpy, "TRUE")

            r = regexp.MustCompile("\\(!\\)")
            statement_cpy = r.ReplaceAllString(statement_cpy, "FALSE")

            r = regexp.MustCompile("@")
            statement_cpy = r.ReplaceAllString(statement_cpy, "TRUE")

            r = regexp.MustCompile("!")
            statement_cpy = r.ReplaceAllString(statement_cpy, "FALSE")

            r = regexp.MustCompile("\\(TRUE\\)")
            statement_cpy = r.ReplaceAllString(statement_cpy, "TRUE")

            r = regexp.MustCompile("\\(FALSE\\)")
            statement_cpy = r.ReplaceAllString(statement_cpy, "FALSE")

            r = regexp.MustCompile("TRUEv(TRUE|FALSE)")
            statement_cpy = r.ReplaceAllString(statement_cpy, "TRUE")

            r = regexp.MustCompile("(TRUE|FALSE)vTRUE")
            statement_cpy = r.ReplaceAllString(statement_cpy, "TRUE")

            r = regexp.MustCompile("FALSEvFALSE")
            statement_cpy = r.ReplaceAllString(statement_cpy, "FALSE")
        }

        //fmt.Printf("new statement: %s\n\n", statement_cpy)

        if statement_cpy == "FALSE" {
            fmt.Println(row)
            new_clause := clause.Clause{}
            for _, l := range row {
                if len(l) == 1 {
                    lit := literal.Literal{Name: l, Negated: true}
                    new_clause.Append(lit)
                }else{
                    lit := literal.Literal{Name: l[2:3], Negated: false}
                    new_clause.Append(lit)
                }

            }
            fmt.Println(new_clause)
            new_clauseset.Append(new_clause)
        }

    }
    fmt.Println("the clauseset:")
    fmt.Println(new_clauseset)
}
