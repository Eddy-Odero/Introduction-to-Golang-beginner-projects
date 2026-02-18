
package main

import (
	"fmt"
	"os"
)

func isBalanced(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		switch ch {

		case '(', '[', '{':
			stack = append(stack, ch)

		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if (ch == ')' && top != '(') ||
				(ch == ']' && top != '[') ||
				(ch == '}' && top != '{') {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	for i := 1; i < len(os.Args); i++ {
		if isBalanced(os.Args[i]) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
