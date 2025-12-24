package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}

func mirrorVowels(s string) string {
	r := []rune(s)

	var vowels []rune
	for _, ch := range r {
		if isVowel(ch) {
			vowels = append(vowels, ch)
		}
	}
	if len(vowels) == 0 {
		return s
	}
	j := len(vowels) - 1
	for i := 0; i < len(r); i++ {
		if isVowel(r[i]) {
			r[i] = vowels[j]
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

	for i, arg := range args {
		result := mirrorVowels(arg)

		for _, ch := range result {
			z01.PrintRune(ch)
		}
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}
}
