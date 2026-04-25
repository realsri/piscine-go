// package main

package piscine

// import (
// 	"fmt"
// 	//"piscine"
// )

func Abort(a, b, c, d, e int) int {
	var res []int
	res = []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4-i; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}
	return res[2]
}

// func main() {
// 	middle := Abort(2, 3, 8, 5, 7)
// 	fmt.Println(middle)
// }
