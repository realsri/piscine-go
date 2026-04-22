package main

import (
	//"github.com/01-edu/z01"
	"bufio"
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			fmt.Print(input)
		}
		return
	}
	for i := 1; i < len(arg); i++ {
		content, err := os.ReadFile(arg[i])
		if err != nil {
			fmt.Printf("ERROR: open %v: No such file or directory", arg[i])
		}
		fmt.Println(string(content))
	}
}
