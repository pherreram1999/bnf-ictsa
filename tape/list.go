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

func insertNode(SymbolCell **CellNode, symbol rune) *CellNode {
	if *SymbolCell == nil {
		*SymbolCell = &CellNode{Symbol: symbol}
		return *SymbolCell
	}
	nav := *SymbolCell
	for nav.Right != nil {
		nav = nav.Right
	}
	nav.Right = &CellNode{Symbol: symbol}
	nav.Right.Left = nav
	return nav.Right
}

func insertFromString(list **CellNode, symbols string) {
	for _, symbol := range symbols {
		insertNode(list, symbol)
	}
}

func getLastNode(list **CellNode) *CellNode {
	nav := *list
	for nav.Right != nil {
		nav = nav.Right
	}
	return nav
}

func listToString(list **CellNode) string {
	nav := *list
	str := ""
	for nav != nil {
		str += string(nav.Symbol)
	}
	return str
}

func (tape *Tape) Add(symbol rune) *CellNode {
	insertNode(&tape.SymbolCell, symbol)
	tape.Length++
	return nil
}

func (tape *Tape) AddFromString(str string) *CellNode {
	var lastNode *CellNode
	for _, char := range str {
		lastNode = tape.Add(char)
	}
	return lastNode
}

// methods

func New() *Tape {
	t := &Tape{}
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
	if t.IFsLen >= t.MaxIFs {
		return // llegamos a los ifs deseados
	}
	fmt.Printf("%s (%s)\n", t, cell)

	if cell.Symbol == 'S' || cell.Symbol == 'A' {
		left := cell.Left
		end := cell.Right
		var newTape *CellNode
		if cell.Symbol == 'S' {
			insertFromString(&newTape, "iCtSA")
			t.IFsLen++
		} else if cell.Symbol == 'A' {
			insertFromString(&newTape, ";eS")
		}

		if left != nil {
			left.Right = newTape
		}
		lastNode := getLastNode(&newTape)
		lastNode.Right = end
		if end != nil {
			end.Left = lastNode
		}
	}

	if t.Rand.Float64() >= .50 && cell.Right != nil {
		// se mueve a la derecha
		t.backusNor(cell.Right)
	} else if cell.Left != nil {
		// se mueve a la izquierda
		t.backusNor(cell.Left)
	} else {
		t.backusNor(cell.Right)
	}

}

func (tape *Tape) MakeProductions(maxIfs int) {
	tape.MaxIFs = maxIfs
	tape.backusNor(tape.SymbolCell)
}
