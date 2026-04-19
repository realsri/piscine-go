// package main
package piscine

/*
import (
	"fmt"
	//"piscine"
)
*/

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arr := make([]int, max-min)

	for i := 0; i < max-min; i++ {
		arr[i] = min + i
	}

	return arr
}

/*
func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
*/
