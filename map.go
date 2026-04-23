package piscine

// package main

// import (
// 	"fmt"
// )

func Map(f func(int) bool, a []int) []bool {
	len := len(a)
	// var res []bool
	res := make([]bool, len)
	for i, c := range a {
		res[i] = f(c)
	}
	return res
}

// func IsPrime1(nb int) bool {
// 	if nb <= 1 {
// 		return false
// 	}
// 	if nb == 2 {
// 		return true
// 	}
// 	if nb%2 == 0 {
// 		return false
// 	}
// 	for i := 3; i*i <= nb; i += 2 {
// 		if nb%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	result := Map(IsPrime1, a)
// 	fmt.Println(result)
// }
