// package main
package piscine

/*
import (
	"fmt"
)*/

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return IterativePower(nb, power-1) * nb
}

/*
func main() {
	fmt.Println(IterativePower(4, 3))
}*/
