package main

import "fmt"

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	digitFound := -1
	for i, j := range s {
		if j >= '0' && j <= '9' {
			if digitFound == -1 {
				digitFound = i
			}
			digit := int(j - '0')
			result = result*10 + digit

		} else if j == '-' && digitFound == -1 {
			sign = -1
		}
	}
	return result * sign
}
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
