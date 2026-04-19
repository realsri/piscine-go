// package main
package piscine

/*
import (
	"fmt"
	//"piscine"
)
*/

func AppendRange(min, max int) []int {
	var arr []int

	/*
		if min >= max {
			return nil
		}
	*/

	for i := min; i < max; i++ {
		arr = append(arr, i)
	}

	return arr
}

/*
func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
	//fmt.Println(AppendRange(5, 5))
}
*/
