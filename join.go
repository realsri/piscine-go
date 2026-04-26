package piscine

// package main

// import (
// 	"fmt"
// 	// "piscine"
// )

func Join(strs []string, sep string) string {
	res := ""
	len := len(strs)
	for i, s := range strs {
		res += s
		if i < len-1 {
			res += sep
		}
	}
	return res
}

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(Join(toConcat, ":"))
// }
