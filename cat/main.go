package main

import (
	"bufio"
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func main() {
	arg := os.Args
	if len(arg) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			// fmt.Print(input)
			Print(input)
		}
		return
	}
	for i := 1; i < len(arg); i++ {
		content, err := os.ReadFile(arg[i])
		if err != nil {
			// fmt.Printf("ERROR: open %v: No such file or directory", arg[i])
			Print("ERROR: open ")
			Print(arg[i])
			Print(": No such file or directory\n")
			os.Exit(1)
		}
		// fmt.Println(string(content))
		Print(string(content))
	}
}
