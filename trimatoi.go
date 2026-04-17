// package main
package piscine

func TrimAtoi(s string) int {
	if len(s) < 1 {
		return 0
	}
	out := 0
	isneg := false
	isdig := false
	for _, i := range s {
		if i >= '0' && i <= '9' {
			out = out*10 + int(i-'0')
			isdig = true
		} else if i == '-' && !isdig {
			isneg = true
		}
	}
	if isneg {
		return -out
	}
	return out
}

/*
func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
*/
