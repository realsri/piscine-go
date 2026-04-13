//package main

package piscine

func BasicAtoi(s string) int {
	//str := []rune(s)
	result := 0
	for _, char := range s {
		digit := int(char - '0')
		result = result*10 + digit
	}
	return result
}

/*
func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
*/
