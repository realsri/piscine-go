package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(x *int, y *int) {
	*x = '*'
	*y = '*' / 2
}

// func printNumber(n int) {
// 	if n >= 10 {
// 		printNumber(n / 10)
// 	}
// 	z01.PrintRune(rune(n%10) + '0')
// }

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func main() {
	points := &point{}

	setPoint(&points.x, &points.y)

	printStr("x = ")
	// printNumber(points.x)
	z01.PrintRune(rune(points.x)/10 + '0')
	z01.PrintRune(rune(points.x)%10 + '0')
	printStr(", y = ")
	// printNumber(points.y)
	z01.PrintRune(rune(points.y)/10 + '0')
	z01.PrintRune(rune(points.y)%10 + '0')
	z01.PrintRune('\n')

	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
