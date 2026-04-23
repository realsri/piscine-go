package piscine

// package main

// import (
// 	"fmt"
// 	// "piscine"
// )

// func IsNumeric1(s string) bool {
// 	if len(s) < 1 {
// 		return false
// 	}
// 	for _, c := range s {
// 		if !(c >= '0' && c <= '9') {
// 			return false
// 		}
// 	}
// 	return true
// }

func Any(f func(string) bool, a []string) bool {
	for _, s := range a {
		if f(s) {
			return true
		}
	}
	return false
}

// func main() {
// 	a1 := []string{"Hello", "how", "are", "you"}
// 	a2 := []string{"This", "is", "4", "you"}

// 	result1 := Any(IsNumeric1, a1)
// 	result2 := Any(IsNumeric1, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
