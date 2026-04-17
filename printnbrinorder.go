// package main
package piscine

// import "piscine"
import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	/*
		var n_arr []int
		i := 0
		for n > 0 {
			n_arr[i] = n / 10
			i++
		}

		for i := 0; i < len(n_arr); i++ {
			z01.PrintRune(n_arr[i])
		}*/
	for i := 0; i <= 9; i++ {
		temp := n
		for temp > 0 {
			if temp%10 == i {
				z01.PrintRune(rune(i + '0'))
			}
			temp /= 10
		}
	}
}

/*
func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
*/
