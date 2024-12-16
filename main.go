package main

import (
	"bnf-ifg/symbolList"
	"fmt"
)

/*
S -> iCtSA
A -> ;eS | epsilon
*/
func main() {
	tape := symbolList.New(5)
	fmt.Println("Tape: ", tape)
	tape.MakeProductions()
}
