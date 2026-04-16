// package main
package piscine

/*
import (
	"github.com/01-edu/z01"
)*/

/*
func NRune(s string, n int) rune {
	srune := []rune(s)
	if n <= len(srune) && n > 0 {
		return srune[n-1]
	} else {
		return 0
	}
}*/

func NRune(s string, n int) rune {
	if n < 1 {
		return 0
	}
	count := 0
	for _, i := range s {
		count++
		if count == n {
			return i
		}
	}
	return 0
}

/*
func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
*/
