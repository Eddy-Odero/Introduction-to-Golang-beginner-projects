package main

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printString("NV")
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	b := len(base)
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	var digits []rune
	for nbr > 0 {
		digits = append(digits, rune(base[nbr%b]))
		nbr /= b
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}