package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args)!=2{
		z01.PrintRune('\n')
		return
	}
	s := os.Args[1]
	word:= false
	first := true
	for i:=0;i < len(s);i++{
		if s[i] !=' '&& s[i] !='\t'{
			if !word && !first{
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(s[i]))
			word = true
			first = false
		}else{
			word= false
		}
	}
	z01.PrintRune('\n')
}