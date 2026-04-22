package main

import (
	"fmt"
	"os"
)

func main() {
	//
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("File name missing")
		os.Exit(0)
	}
	if len(arg) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}

	content, _ := os.ReadFile(arg[1])
	fmt.Print(string(content))
}
