package main

import (
	"fmt"
	"os"
)

// func substrcheck(s, sub string) bool {
// 	for i := 0; i < len(s); i++ {
// 		if i+len(sub) <= len(s) && s[i:i+len(sub)] == sub {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	args := os.Args
	len := len(args)

	for i := 1; i < len; i++ {
		if args[i] == "01" || args[i] == "galaxy" || args[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			// os.Exit(0)
			return
			// fmt.Println("\n")
		}
	}
}
