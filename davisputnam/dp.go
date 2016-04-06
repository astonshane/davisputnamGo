package main

import (
	"bufio"
	"davisputnam/literal"
	"davisputnam/clause"
	"davisputnam/clauseset"
	//"davisputnam/connector"
	"fmt"
	"log"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
	"regexp"
)

/*
def Satisfiable(CS):
    if CS = {}:
        return true
    if {} in CS:
        return false
    if {L} in CS:
        return Satisfiable(CS_L)
    select L in lit(CS)
        return Satisfiable(CS_L) | Satisfiable(CS_L')
*/

//Satisfiable implements above function
func Satisfiable(CS clauseset.ClauseSet) bool {
	fmt.Printf("\n%sSatisfiable(%s)\n", strings.Repeat(" ", CS.Indent), CS)
	CS.Indent += 2

	//if CS = {} : return true
	if CS.Len() == 0 {
		fmt.Printf("%sEmpty ClauseSet, returning true\n", strings.Repeat(" ", CS.Indent))
		return true
	}

	//if {} in CS : return false
	firstElement, err := CS.FirstElement()
	if err != nil {
		log.Fatal(err)
	}
	if clause.Len(firstElement) == 0 {
		fmt.Printf("%s{} found in ClauseSet, returning false\n", strings.Repeat(" ", CS.Indent))
		return false
	}

	//select L in lit(CS): return Satisfiable(CS_L) || Satisfiable(CS_L')
	nextLiteral, err2 := CS.NextLiteral()
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("%sSpliting on %s\n", strings.Repeat(" ", CS.Indent), nextLiteral)

	CSL := CS.Reduce(nextLiteral)
	CSL.Indent = CS.Indent
	CSR := CS.Reduce(nextLiteral.Negation())
	CSR.Indent = CS.Indent

	return Satisfiable(CSL) || Satisfiable(CSR)
}

//FindValidity finds the validity of the argument (given as a ClauseSet)
func FindValidity(CS clauseset.ClauseSet) {
	fmt.Printf("Starting ClauseSet: %s\n", CS)

	CS.Indent = 0
	sat := Satisfiable(CS)

	fmt.Print("Conclusion: ")
	if sat {
		fmt.Printf("Satisfiable(%s) == %t; INVALID\n", CS, sat)
	} else {
		fmt.Printf("Satisfiable(%s) == %t; VALID\n", CS, sat)
	}
}

func clean(sentence string) string {
	result := strings.Replace(sentence, "<->", "+xnor+", -1)
	result = strings.Replace(result, "->", "+implies+", -1)
	result = strings.Replace(result, "v", "+or+", -1)
	result = strings.Replace(result, "^", "+and+", -1)
	result = strings.Replace(result, "~", "+not+", -1)
	return result
}

func findMatch(cnf string) clauseset.ClauseSet {
	cs := clauseset.ClauseSet{}
	clauses := strings.Split(cnf, "AND")
	for _, c := range clauses{
		cl := clause.Clause{}
		c = strings.Trim(strings.TrimSpace(c), "()")
		lits := strings.Split(c, "OR")
		for _, lit := range lits{
			lit = strings.Trim(strings.TrimSpace(lit), "()")
			if len(lit) == 1{
				l := literal.Literal{Name: lit, Negated: false}
				cl.Append(l)
			}else{
				l := literal.Literal{Name: strings.Split(lit, " ")[1], Negated: true}
				cl.Append(l)
			}


		}
		cs.Append(cl)
	}
	return cs
}

func getCNF(sentence string) clauseset.ClauseSet {
	cs := clauseset.ClauseSet{}
	url := fmt.Sprintf("http://api.wolframalpha.com/v1/query?input=BooleanConvert[%s,%%22CNF%%22]&appid=2Y4TEV-W2AETK4T5K", clean(sentence))
	response, err := http.Get(url)
	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
		lines := strings.Split(string(contents), "\n")
		found_first := false

		for _, line := range lines {
			if strings.Contains(line, "<plaintext>") {
				if !found_first {
					found_first = true
				}else{

					re := regexp.MustCompile("<plaintext>(.*)</plaintext>")
					cnf := re.FindStringSubmatch(line)[1]

					newCS := findMatch(cnf)
					cs = clauseset.Combine(cs, newCS)
					break
				}
			}
		}
    }

	return cs
}

//ConstructCS reads in premisies from a file and returns its ClauseSet
func ConstructCS(file string) clauseset.ClauseSet {
	newCS := clauseset.ClauseSet{}

	infile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer infile.Close()
	scanner := bufio.NewScanner(infile)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	contin := true
	for contin {
		line := scanner.Text()

		if !scanner.Scan() {
			line = fmt.Sprintf("~(%s)", line)
			contin = false
		}

		cs := getCNF(line)
		newCS = clauseset.Combine(newCS, cs)
	}

	return newCS
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Useage: go dp.go <input file>")
		os.Exit(0)
	}
	CS := ConstructCS(os.Args[1])
	fmt.Println(CS)
	FindValidity(CS)

}
