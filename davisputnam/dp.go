package main

import (
	"bufio"
	"davisputnam/clause"
	"davisputnam/clauseset"
	"davisputnam/cnf"
	"fmt"
	"log"
	"os"
	"strings"
	"encoding/json"
	"os/exec"
)


type Tree struct {
	Name 		string
	Split 		string
	Children	[]Tree
}
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
func Satisfiable(CS clauseset.ClauseSet) (bool, Tree) {
	tree := Tree{Name: CS.String(), Split: ""}
	fmt.Printf("\n%sSatisfiable(%s)\n", strings.Repeat(" ", CS.Indent), CS)
	CS.Indent += 2

	//if CS = {} : return true
	if CS.Len() == 0 {
		fmt.Printf("%sEmpty ClauseSet, returning true\n", strings.Repeat(" ", CS.Indent))
		openTree := Tree{Name: "O", Split: ""}
		tree.Children = append(tree.Children, openTree)
		return true, tree
	}

	//if {} in CS : return false
	firstElement, err := CS.FirstElement()
	if err != nil {
		log.Fatal(err)
	}
	if clause.Len(firstElement) == 0 {
		fmt.Printf("%s{} found in ClauseSet, returning false\n", strings.Repeat(" ", CS.Indent))
		closedTree := Tree{Name: "X", Split: ""}
		tree.Children = append(tree.Children, closedTree)
		return false, tree
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

	left, tree_l := Satisfiable(CSL)
	right, tree_r := Satisfiable(CSR)

	tree_l.Split = nextLiteral.String()
	tree_r.Split = nextLiteral.Negation().String()

	tree.Children = append(tree.Children, tree_l)
	tree.Children = append(tree.Children, tree_r)
	return left || right, tree
}

//FindValidity finds the validity of the argument (given as a ClauseSet)
func FindValidity(CS clauseset.ClauseSet, filename string) {
	fmt.Printf("Starting ClauseSet: %s\n", CS)

	CS.Indent = 0
	sat, tree := Satisfiable(CS)
	treeJson, _ := json.Marshal(tree)

	f, _ := os.Create(filename)
	f.WriteString(string(treeJson))

	fmt.Print("Conclusion: ")
	if sat {
		fmt.Printf("Satisfiable(%s) == %t; INVALID\n", CS, sat)
	} else {
		fmt.Printf("Satisfiable(%s) == %t; VALID\n", CS, sat)
	}
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
		fmt.Println(line)

		cs := cnf.ToCNF(line)
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
	filename := os.Args[1][:len(os.Args[1])-4]

	FindValidity(CS, filename+".json")

	cmd := "python"
	args := []string{"graph.py", filename+".json"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
