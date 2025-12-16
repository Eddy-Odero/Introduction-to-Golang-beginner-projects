package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]

	if len(arg) >= 2 && arg[0] == '.' && arg[1] == '/' {
		arg = arg[2:]
	}

	for _, r := range arg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
