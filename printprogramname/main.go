package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	arg0 := args[0]

	len := len(arg0)

	for i := 2; i < len; i++ {
		z01.PrintRune(rune(arg0[i]))
	}
	z01.PrintRune('\n')
}
