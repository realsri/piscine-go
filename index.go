package main

//package piscine
/*
import (
	"fmt"
	//"piscine"
)
*/

func Index(s string, toFind string) int {
	lenF := len(toFind)
	lenS := len(s)

	if lenF == 0 {
		return 0
	}

	for i := 0; i <= lenS-lenF; i++ {
		if toFind == s[i:i+lenF] {
			return i
		}
	}
	return -1
}

/*
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}*/
