package main

import (
	"strconv"
)

func isPiece(p string) bool {
	i := int(p[0])
	if (64 < i) && (i < 91) || (96 < i) && (i < 123) {
		return true
	}
	return false
}

func isInt(p string) bool {
	_, err := strconv.Atoi(p)
	if err != nil {
		return false
	}
	return true
}

