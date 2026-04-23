// package main
package piscine

/*
import (

	"piscine"

	"github.com/01-edu/z01"

)
*/
func ForEach(f func(int), a []int) {
	for _, c := range a {
		f(c)
	}
}

/*
func PrintNbr1(n int) {
	z01.PrintRune(rune(n) + '0')
}
*/
/*
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(piscine.PrintNbr, a)
}
*/
