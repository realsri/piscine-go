// package main

package piscine

// import (
// 	"fmt"
// 	// "piscine"
// )

func Max(a []int) int {
	max := a[0]
	for _, n := range a {
		if n > max {
			max = n
		}
	}
	return max
}

// func main() {
// 	a := []int{23, 123, 1, 11, 55, 93}
// 	max := Max(a)
// 	fmt.Println(max)
// }
