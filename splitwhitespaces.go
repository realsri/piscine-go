// package main
package piscine

/*
import (
	"fmt"
	// "piscine"
)
*/

func SplitWhiteSpaces(s string) []string {
	var res []string
	var w string

	for _, c := range s {
		if c == ' ' || c == '\t' || c == '\n' {
			if w != "" {
				res = append(res, w)
				w = ""
			}
		} else {
			w += string(c)
		}
	}
	return res
}

/*
func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
*/
