package main

import (
	"fmt"
	"os"
)

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}
func cleanSpaces(s string)string{
	var result []rune
	inword := false
	for _, r := range s{
		if !isSpace (r){
			if !inword && len(result) > 0 {
				result = append(result, ' ')
			}
			result = append(result, r)
			inword = true
		}else{
			inword = false
		}
	}
	return string(result)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	input := os.Args[1]
	cleaned := cleanSpaces(input)
	fmt.Println(cleaned)
}
