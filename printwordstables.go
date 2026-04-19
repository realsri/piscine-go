package piscine

import "github.com/01-edu/z01"

// package main

// import (
// 	//"piscine"

// 	"github.com/01-edu/z01"
// )

func PrintWordsTables(a []string) {
	for _, w := range a {
		for _, c := range w {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}

// func SplitWhiteSpaces1(s string) []string {
// 	var res []string
// 	var w string

// 	for _, c := range s {
// 		if c == ' ' || c == '\t' || c == '\n' {
// 			if w != "" {
// 				res = append(res, w)
// 				w = ""
// 			}
// 		} else {
// 			w += string(c)
// 		}
// 	}

// 	if w != "" {
// 		res = append(res, w)
// 	}
// 	return res
// }

// func main() {
// 	a := SplitWhiteSpaces1("Hello how are you?")
// 	PrintWordsTables(a)
// }
