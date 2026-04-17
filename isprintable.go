// package main
package piscine

/*
import (
	"fmt"
	//"piscine"
)*/

func IsPrintable(s string) bool {
	if len(s) < 1 {
		return false
	}
	for _, c := range s {
		if !(c >= 32 && c <= 126) {
			return false
		}
	}
	return true
}

/*
func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}
*/
