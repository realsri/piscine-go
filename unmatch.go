package piscine

// package main

// import (
// 	"fmt"
// 	// "piscine"
// )

func Unmatch(a []int) int {
	len := len(a)
	for i := 0; i < len; i++ {
		count := 0
		for j := 0; j < len; j++ {
			if a[i] == a[j] {
				count++
			}
		}
		if count%2 != 0 {
			return a[i]
		}
	}
	return -1
}

// func main() {
// 	a := []int{1, 2, 3, 1, 2, 3, 4}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
