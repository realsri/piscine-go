// package main

// import "fmt"
package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb < 0 {
		return 0
	} else {
		return IterativeFactorial(nb-1) * nb
	}
}

/*
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
*/
