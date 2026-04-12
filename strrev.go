// package main
package piscine

/*
import (
	"fmt"
)*/

func StrRev(s string) string {
	changeStr := []byte(s)
	len := len(changeStr)
	var temp byte
	for i := 0; i < len/2; i++ {
		temp = changeStr[i]
		changeStr[i] = changeStr[len-i-1]
		changeStr[len-i-1] = temp
	}
	return string(changeStr)
}

/*
func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
*/
