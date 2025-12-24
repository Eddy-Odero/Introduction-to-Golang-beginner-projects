package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	for _, v := range "aeiouAEIOU" {
		if r == v {
			return true
		}
	}
	return false
}

func rotateVowels(s string) string {
	r := []rune(s)
	var v []rune

	for _, c := range r {
		if isVowel(c) {
			v = append(v, c)
		}
	}

	j := len(v) - 1
	for i := range r {
		if isVowel(r[i]) {
			r[i] = v[j]
			j--
		}
	}
	return string(r)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	for i, a := range args {
		for _, c := range rotateVowels(a) {
			z01.PrintRune(c)
		}
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}
}
