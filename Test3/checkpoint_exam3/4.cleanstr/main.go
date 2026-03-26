package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
	z01.PrintRune('\n')
		return
	}

	s := os.Args[1]
	inword := false
	firstWord := true

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' {
				if !inword && !firstWord{
					z01.PrintRune(' ')
				}
				z01.PrintRune(rune(s[i]))
				firstWord = false
				inword = true
		} else {
			inword = false
		}
	}
	z01.PrintRune('\n')
	}