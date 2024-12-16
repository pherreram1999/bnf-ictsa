package symbolList

import (
	"math/rand"
	"time"
)

type (
	SymbolNode struct {
		Symbol rune
		Right  *SymbolNode
		Left   *SymbolNode
	}

	Tape struct {
		SymbolNode *SymbolNode // es nuestro puntero de lista doblemente enlazada
		Head       *SymbolNode // donde se encuentra el cabezal donde se turing
		Length     int
		IFsLen     int
		MaxIFs     int
		Rand       *rand.Rand
	}
)

// static functions
func (tape *Tape) Add(symbol rune) *SymbolNode {
	if tape.SymbolNode == nil {
		tape.SymbolNode = &SymbolNode{Symbol: symbol}
		return tape.SymbolNode
	}
	nav := tape.SymbolNode
	for nav.Right != nil {
		nav = nav.Right
	}
	nav.Right = &SymbolNode{Symbol: symbol}
	nav.Right.Left = nav
	tape.Length++
	return nav.Right
}

func (tape *Tape) AddFromString(str string) *SymbolNode {
	var lastNode *SymbolNode
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
	nav := tape.SymbolNode
	for nav != nil {
		str += string(nav.Symbol)
		nav = nav.Right
	}
	return str
}

func backusNor() {

}

func (tape *Tape) MakeProductions() {

	backusNor()

}
