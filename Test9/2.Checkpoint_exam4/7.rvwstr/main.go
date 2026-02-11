package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	n := len(s)

	end := n

	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			fmt.Print(s[i+1 : end])
			fmt.Print(" ")
			end = i
		}
	}

	
	fmt.Print(s[0:end])
	fmt.Println()
}
