package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	n := []rune(s)

	return n[len(n)-1]
}
func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
