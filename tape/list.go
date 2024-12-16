package tape

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	CellNode struct {
		Symbol rune
		Right  *CellNode
		Left   *CellNode
	}

	Tape struct {
		SymbolCell *CellNode // es nuestro puntero de lista doblemente enlazada
		Head       *CellNode // donde se encuentra el cabezal donde se turing
		Length     int
		IFsLen     int
		MaxIFs     int
		Rand       *rand.Rand
	}
)

// static functions
func (tape *Tape) Add(symbol rune) *CellNode {
	if tape.SymbolCell == nil {
		tape.SymbolCell = &CellNode{Symbol: symbol}
		return tape.SymbolCell
	}
	nav := tape.SymbolCell
	for nav.Right != nil {
		nav = nav.Right
	}
	nav.Right = &CellNode{Symbol: symbol}
	nav.Right.Left = nav
	tape.Length++
	return nav.Right
}

func (tape *Tape) AddFromString(str string) *CellNode {
	var lastNode *CellNode
	for _, char := range str {
		lastNode = tape.Add(char)
	}
	return lastNode
}

// methods

func New(maxIFs int) *Tape {
	t := &Tape{MaxIFs: maxIFs}
	t.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	t.AddFromString("iCtSA")
	return t
}

func (tape *Tape) String() string {
	str := ""
	nav := tape.SymbolCell
	for nav != nil {
		str += string(nav.Symbol)
		nav = nav.Right
	}
	return str
}

func (sc *CellNode) String() string {
	return string(sc.Symbol)
}

func (t *Tape) backusNor(cell *CellNode) {
	desition := t.Rand.Float64()
	fmt.Println(cell, desition)
	time.Sleep(time.Second)
	if desition >= .50 && cell.Right != nil {
		// se mueve a la derecha
		t.backusNor(cell.Right)
	} else if cell.Left != nil {
		// se mueve a la izquierda
		t.backusNor(cell.Left)
	}

}

func (tape *Tape) MakeProductions() {
	tape.backusNor(tape.SymbolCell)
}
