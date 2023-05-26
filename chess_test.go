package main

import "testing"
import "fmt"

func TestInit(t *testing.T) {
	// x:
	letters := func() string {
		var s string
		for i := 'a'; i <= 'h'; i++ {
			s += string(i)
		}
		return s
	}()
	// y:
	digits := func() string {
		var s string
		for i := '1'; i <= '8'; i++ {
			s += string(i)
		}
		return s
	}()
	coords := letters + digits
	x := coords[:len(letters)]
	y := coords[len(letters):]
	func(strings ...string) {
		for _, v := range strings {
			fmt.Println(v)
		}
	}(letters, digits, coords, x, y)

}
