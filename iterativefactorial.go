// package main
package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 {
		return IterativeFactorial(nb-1) * nb
	} else {
		return 1
	}
}

/*
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}*/
