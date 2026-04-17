// package main
package piscine

/*
import (
	"fmt"
	//"piscine"
)*/

func ToUpper(s string) string {
	us := []rune(s)
	for i, c := range s {
		if c >= 'a' && c <= 'z' {
			us[i] = c - 32
		}
	}
	return string(us)
}

/*
func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}*/
