package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := a; b <= '9'; b++ {
			for i := '0'; i <= '9'; i++ {
				for j := i; j <= '9'; j++ {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(i)
					z01.PrintRune(j)
					if i != '8' { // comma space not needed at the last
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
