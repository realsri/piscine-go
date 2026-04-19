package main

// package piscine

import (
	"fmt"
	//"piscine"
)

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := 0
	for i := 0; i < len(nbr); i++ {
		val := 0
		for j := 0; j < len(baseFrom); j++ {
			if nbr[i] == baseFrom[j] {
				val = j
				break
			}
		}
		n = n*len(baseFrom) + val
	}

	if n == 0 {
		return string(baseTo[0])
	}

	res := ""
	for n > 0 {
		res = string(baseTo[n%len(baseTo)]) + res
		n /= len(baseTo)
	}
	return res
}

func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}
