// package main
package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb > 127 {
		return false
	}
	for i := 2; i < nb; i++ {
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
}
*/
