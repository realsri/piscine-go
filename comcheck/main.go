package main

import (
	"fmt"
	"os"
)

func substrcheck(s, sub string) bool {
	for i := 0; i < len(s); i++ {
		if i+len(sub) <= len(s) && s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args
	len := len(args)

	for i := 1; i < len; i++ {
		if substrcheck(args[i], "01") || substrcheck(args[i], "galaxy") {
			fmt.Println("Alert!!!")
			// fmt.Println("\n")
		}
	}
}
