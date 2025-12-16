package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	for _, i := range args {
		z01.PrintRune(i)
	}

}
