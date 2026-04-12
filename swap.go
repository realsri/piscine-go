// package main
package piscine

// import "fmt"
func Swap(a *int, b *int) {
	*a, *b = *b, *a
	//temp := *a
	//*a = *b
	//*b = temp
}

/*
func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
*/
