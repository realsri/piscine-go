// package main
package piscine

/*
import (
	"fmt"
	//"piscine"
)*/

func ToLower(s string) string {
	us := []rune(s)
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			us[i] = c + 32
		}
	}
	return string(us)
}

/*
func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}*/
