package main

import "fmt"

func toBase10(s string, base string) int {
	result := 0
	baseLen := len(base)
	for _, i := range s {
		digit := 0
		for a, b := range base {
			if i == b {
				digit = a
				break
			}
		}
		result = result*baseLen + digit
	}
	return result
}

func fromBase10(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}
	baseLen := len(base)
	result := ""
	for n > 0 {
		rem := n % baseLen
		result = string(base[rem]) + result
		n /= baseLen
	}
	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	value := toBase10(nbr, baseFrom)
	return fromBase10(value, baseTo)
}

func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}
