package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args

	lena := len(arg)

	for i := 0; i < lena; i++ {
		argrune := []rune(arg[i])
		for j := 0; j < len(argrune); j++ {
			z01.PrintRune(argrune[j])
		}
		z01.PrintRune('\n')
	}
}
