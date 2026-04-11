package piscine

// package main

// import "fmt"
func UltimateDivMod(a *int, b *int) {
	tmp := *a / *b
	*b = *a % *b
	*a = tmp
}

/*
func main() {
	a := 23
	b := 2

	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Print(b)
}*/
