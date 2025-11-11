package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	n := []rune(s)

	return n[0]
}
func main() {
	z01.PrintRune(FirstRune("Hello!"))
}
