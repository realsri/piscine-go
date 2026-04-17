// package main
package piscine

/*
import (
	"fmt"
	//"piscine"
)
*/

func IsUpper(s string) bool {
	if len(s) < 1 {
		return false
	}
	for _, c := range s {
		if !(c >= 'A' && c <= 'Z') {
			return false
		}
	}
	return true
}

/*
func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}*/
