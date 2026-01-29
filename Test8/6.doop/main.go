
package main

import (
	"fmt"
	"os"
)

func atoi(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}

	sign := int64(1)
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	if i == len(s) {
		return 0, false
	}

	var n int64
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		d := int64(s[i] - '0')

		// overflow check
		if n > (1<<63-1-d)/10 {
			return 0, false
		}
		n = n*10 + d
	}

	n *= sign

	if n > (1<<63-1) || n < -(1<<63) {
		return 0, false
	}

	return n, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoi(os.Args[1])
	op := os.Args[2]
	b, ok2 := atoi(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	if op == "/" && b == 0 {
		fmt.Println("No division by 0")
		return
	}

	if op == "%" && b == 0 {
		fmt.Println("No modulo by 0")
		return
	}

	var res int64
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "%":
		res = a % b
	default:
		return
	}

	// overflow check after operation
	if res > (1<<63-1) || res < -(1<<63) {
		return
	}

	fmt.Println(res)
}
