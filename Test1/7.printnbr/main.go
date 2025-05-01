package main

import "github.com/01-edu/z01"

//undestand that this problem want  us to implement itoa function
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	var digit []int
	for n > 0 {
		digit = append(digit, n%10)
		n /= 10
	}
	for i := len(digit) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digit[i] + '0'))
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
