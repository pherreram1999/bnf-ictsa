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
	t := tape.New()
	fmt.Println("Tape: ", t)
	t.MakeProductions(5)
}
