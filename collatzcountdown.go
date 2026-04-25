// package main

package piscine

// import (
// 	"fmt"
// 	// "piscine"
// )

func CollatzCountdown(start int) int {
	res, count := start, 0
	if start <= 0 {
		return -1
	}
	for i := 0; res != 1; i++ {
		if res%2 == 0 { // even
			res = res / 2
		} else { // odd
			res = 3*res + 1
		}
		count++
	}
	return count
}

// func main() {
// 	steps := CollatzCountdown(12)
// 	fmt.Println(steps)
// }
