// package main

// import "fmt"
package piscine

/*
	func Sqrt1(nb int) int {
		if nb == 0 {
			return 0
		}
		for i := 0; i <= nb; i++ {
			if nb == i*i {
				return i
			}
		}
		return 0
	}
*/
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb%2 == 0 {
		return false
	}
	for i := 2; i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

/*
func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}*/
