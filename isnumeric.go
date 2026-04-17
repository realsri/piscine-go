// package main
package piscine

/*
import (
	"fmt"
	//"piscine"
)*/

func IsNumeric(s string) bool {
	if len(s) < 1 {
		return false
	}
	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}

/*
func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}*/
