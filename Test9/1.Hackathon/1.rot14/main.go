package main

import (
	

	"github.com/01-edu/z01"
)
func Rot14(s string) string {
	result := []rune(s)

	for i, c := range result {
		if c >= 'a' && c <= 'z' {
			result[i] = ((c-'a'+14)%26) + 'a'
		} else if c >= 'A' && c <= 'Z' {
			result[i] =  ((c-'A'+14)%26)+ 'A' 
		}
	}
	return string(result)
}


func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
