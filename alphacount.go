// package main
package piscine

/*
import (

	"fmt"
	//"piscine"

)
*/
func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if char >= 'A' && char <= 'z' {
			count++
		}
	}
	return count
}

/*
func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
*/
