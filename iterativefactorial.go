// package main
package piscine

/*
func IterativeFactorial(nb int) int { //oops! it is recursive factorial
	if nb == 0 {
		return 1
	} else if nb < 0 {
		return 0
	} else {
		return IterativeFactorial(nb-1) * nb
	}
}*/

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		if result > 0 && result*i < result {
			return 0
		}
		result = result * i
	}
	return result
}

/*

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
*/
