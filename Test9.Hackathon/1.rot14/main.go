package main

import (
	

	"github.com/01-edu/z01"
)
func Rot14(s string) string {

}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
