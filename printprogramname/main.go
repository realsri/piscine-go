package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	arg0 := args[0]

	arg0rune := []rune(arg0)

	len := len(arg0rune)

	for i := 2; i < len; i++ {
		z01.PrintRune(rune(arg0rune[i]))
	}
	z01.PrintRune('\n')
}
