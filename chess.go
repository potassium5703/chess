package main

import (
	"fmt"
	"strings"
)

var x, y string

func init() {
	var (
		// X:
		letters = func() string {
			var s string
			for i := 'a'; i <= 'h'; i++ {
				s += string(i)
			}
			return s
		}()
		// Y:
		digits = func() string {
			var s string
			for i := '1'; i <= '8'; i++ {
				s += string(i)
			}
			return s
		}()

		coords = letters + digits
	)
	x = coords[:len(letters)]
	y = coords[len(letters):]

}

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
			fmt.Printf("%d ", reverseY(y))
		}
		for x, _ := range b {
			fmt.Printf("%s ", string(b[x][y]))
		}
		fmt.Print("\n")
	}
	if withcoords {
		fmt.Printf(" ")
		for char := 'a'; char <= 'h'; char++ {
			fmt.Printf(" %s", string(char))
		}
		fmt.Printf("\n")
	}
}

// converts Position string value to coords on Board
// func (p Position) ToCoords() (int, int) {
// 	i := strings.Index(coords)
// }

// Moves piece
// For exmaple: "d2", "d4" moves piece from d2 to d4
func (b *Board) Move(o, n Position) {
	type Coords struct {
		X, Y int
	}
	var old, new Coords
	old.X = strings.Index(x, string(o[0]))
	if old.X < 0 {
		fmt.Errorf("invalid position")
	}
	old.Y = strings.Index(y, string(o[1]))
	if old.Y < 0 {
		fmt.Errorf("invalid position")
	}
	old.Y = reverseY(old.Y) - 1

	new.X = strings.Index(x, string(n[0]))
	if new.X < 0 {
		fmt.Errorf("invalid position")
	}
	new.Y = strings.Index(y, string(n[1]))
	if new.Y < 0 {
		fmt.Errorf("invalid position")
	}
	new.Y = reverseY(new.Y) - 1

	b[old.X][old.Y], b[new.X][new.Y] = b[new.X][new.Y], b[old.X][old.Y]
}

func main() {
	board := initialPosition.Parse()
	board.Move("d2", "d4")
	board.Print(true)
}
