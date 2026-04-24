package piscine

// package main

// import (
// 	"fmt"
// 	//"piscine"
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

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, s := range tab {
		if f(s) {
			count++
		}
	}
	return count
}

// func main() {
// 	tab1 := []string{"Hello", "how", "are", "you"}
// 	tab2 := []string{"This", "1", "is", "4", "you"}
// 	answer1 := CountIf(IsNumeric1, tab1)
// 	answer2 := CountIf(IsNumeric1, tab2)
// 	fmt.Println(answer1)
// 	fmt.Println(answer2)
// }
