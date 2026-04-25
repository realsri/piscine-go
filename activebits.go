// package main

package piscine

// import (
// 	"fmt"
// 	// "piscine"
// )

func ActiveBits(n int) int {
	// return n / 2
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

// func main() {
// 	fmt.Println(ActiveBits(7))
// }
