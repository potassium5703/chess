package main

import (
	"strconv"
)

// https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
type FEN string

func (f FEN) Parse() Board {
	// строка p может содержать:
	// шахматная фигура
	// переход на новую строку
	// клеточки которые надо пропустить
	// пустые ячейки

	var b Board
	var x, y int

	for _, v := range f {
		i := string(v)
		switch true {
		case isPiece(i):
			b[x][y] = byte(i[0])
			if x < 7 {
				x++
			}
		case isInt(i):
			add, _ := strconv.Atoi(i)
			add--
			for add >= 0 {
				b[add][y] = '.'
				add = add - 1
			}
		case i == "/":
			y++
			x = 0
		case i == " ":
			return b
		default:
			return b
		}
	}
	return b
}
