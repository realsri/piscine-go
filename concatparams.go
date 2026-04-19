// package main
package piscine

/*
import (

	"fmt"
	//"piscine"

)
*/
func ConcatParams(args []string) string {
	var res string

	for i, arg := range args {
		res += arg
		if i < len(args)-1 {
			res += "\n"
		}
	}

	return res
}

/*
func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
*/
