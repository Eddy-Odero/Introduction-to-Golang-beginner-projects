package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	isUpper := false
	if args[0] == "--upper" {
		isUpper = true
		args = args[1:]
	}

	for _, arg := range args {
		n := 0
		isValid := true
		if len(arg) == 0 {
			isValid = false
		}
		for i := 0; i < len(arg); i++ {
			if arg[i] < '0' || arg[i] > '9' {
				isValid = false
				break
			}
			n = n*10 + int(arg[i]-'0')
		}
		if isValid && n >= 1 && n <= 26 {
			var r rune
			if isUpper {
				r = rune('A' + n - 1)
			} else {
				r = rune('a' + n - 1)
			}
			z01.PrintRune(r)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
