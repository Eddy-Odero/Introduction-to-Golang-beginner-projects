package main

import (
	"fmt"
	"os"
)

func splitArrayContent(s string) []string {
	var result []string
	current := ""

	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			result = append(result, current)
			current = ""
		} else if s[i] != ' ' {
			current += string(s[i])
		}
	}

	result = append(result, current)
	return result
}

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
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]
	if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	content := arrStr[1 : len(arrStr)-1]
	elements := splitArrayContent(content)

	var numbers []int
	for _, el := range elements {
		num, ok := atoiStrict(el)
		if !ok {
			fmt.Printf("Invalid number: %s\n", el)
			return
		}
		numbers = append(numbers, num)
	}

	target, ok := atoiStrict(os.Args[2])
	if !ok {
		fmt.Println("Invalid target sum.")
		return
	}

	pairs := [][]int{}
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
}
