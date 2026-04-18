package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		return
	}

	isUpper := false
	if arg[0] == "--upper" {
		isUpper = true
		arg = arg[1:]
	}

	lena := len(arg)

	for i := 0; i < lena; i++ {
		str := arg[i]
		n := 0
		/*for j := 0; j < len(argrune); j++ {
			z01.PrintRune(argrune[j])
		}*/

		if len(str) == 1 && str[0] >= '0' && str[0] <= '9' {
			n = int(str[0] - '0')
		} else if len(str) == 2 && str[0] >= '0' && str[0] <= '9' && str[1] >= '0' && str[1] <= '9' {
			n = int(str[0]-'0')*10 + int(str[1]-'0')
		} else {
			z01.PrintRune(' ')
			// z01.PrintRune('\n')
			continue
		}

		if len(str) > 2 {
			z01.PrintRune(' ')
		} else if n >= 1 && n <= 26 {
			if isUpper {
				toprint := n + 64
				z01.PrintRune(rune(toprint))
			} else {
				toprint := n + 96
				z01.PrintRune(rune(toprint))
			}
		} else {
			z01.PrintRune(' ')
		}
		// z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
