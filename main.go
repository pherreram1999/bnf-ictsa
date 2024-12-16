package main

import (
	"bnf-ifg/tape"
	"fmt"
)

/*
S -> iCtSA
A -> ;eS | epsilon
*/
func main() {
	t := tape.New(5)
	fmt.Println("Tape: ", t)
	t.MakeProductions()
}
