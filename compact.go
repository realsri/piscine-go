// package main
package piscine

// import (
// 	"fmt"
// 	// "piscine"
// )

// const N = 6

func Compact(ptr *[]string) int {
	j := 0
	for i, s := range *ptr {
		if s != "" {
			(*ptr)[j] = (*ptr)[i]
			j++
		}
	}
	*ptr = (*ptr)[:j]
	return j
}

// func main() {
// 	a := make([]string, N)
// 	a[0] = "a"
// 	a[2] = "b"
// 	a[4] = "c"

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}

// 	fmt.Println("Size after compacting:", Compact(&a))

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}
// }
