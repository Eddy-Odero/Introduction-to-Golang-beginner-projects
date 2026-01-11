package main

import (
	"fmt"
	"os"
	"strings"
)

func atoiStrict(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := 1
	i := 0
	result := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	if i == len(s) {
		return 0, false
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		result = result*10 + int(s[i]-'0')
	}

	return result * sign, true
}

func main() {
	// LEVEL 1: Argument validation
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	// LEVEL 2: Array format validation
	arrStr := os.Args[1]
	if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	// LEVEL 3: Split array content
	content := arrStr[1 : len(arrStr)-1]
	elements := strings.Split(content, ",")

	// LEVEL 4–5: Parse numbers
	var numbers []int
	for _, el := range elements {
		el = strings.TrimSpace(el)

		num, ok := atoiStrict(el)
		if !ok {
			fmt.Printf("Invalid number: %s\n", el)
			return
		}
		numbers = append(numbers, num)
	}

	// LEVEL 6: Parse target
	target, ok := atoiStrict(os.Args[2])
	if !ok {
		fmt.Println("Invalid target sum.")
		return
	}

	// LEVEL 7: Find pairs
	pairs := [][]int{}
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	// LEVEL 8: Output
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
}
