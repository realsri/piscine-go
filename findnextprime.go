// package main
package piscine

// import "fmt"
func IsPrime1(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for i := nb; true; i++ {
		if IsPrime1(i) {
			return i
		}
	}
	return 0
}

/*
func main() {
	fmt.Println(FindNextPrime(-1))
	fmt.Println(FindNextPrime(2))
}*/
