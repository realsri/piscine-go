// package main
package piscine

/*
import (
	"fmt"
	//"piscine"
)
*/

func IsLower(s string) bool {
	if len(s) < 1 {
		return false
	}
	for _, c := range s {
		if !(c >= 'a' && c <= 'z') {
			return false
		}
	}
	return true
}

/*
func main() {
	fmt.Println(IsLower("HELLO"))
	fmt.Println(IsLower("HELLO!"))
}*/
