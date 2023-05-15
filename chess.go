package main

import (
	"fmt"
	// "strings"
)

var initialPosition FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Parser interface {
	Parse()
}

type Piece interface {
	LegalMoves() []Position 
}

type Board [8][8]byte
type Position string

func (b Board) Print(withcoords bool) {
	for y, _ := range b {
		if withcoords {
			fmt.Printf("%d ", (y-8)*(-1))
		}
		for x, _ := range b {
			fmt.Printf("%s ", string(b[x][y]))
		}
		fmt.Print("\n")
	}
	if withcoords {
		fmt.Printf(" ")
		for char := 97; char < 105; char++ {
			fmt.Printf(" %s", string(char))
		}
		fmt.Printf("\n")
	}
}

// converts Position string value to coords on Board 
// func (p Position) ToCoords() (int, int) {
// 	i := strings.Index(coords)	
// }

// moves piece 
func (b Board) Move(o, n Position) {
	
}

func main() {
	board := initialPosition.Parse()
	board.Print(true)
}
