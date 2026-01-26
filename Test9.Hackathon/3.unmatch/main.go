package main

import (
	"fmt"
	
)

func Unmatch(a []int) int {
	result := 0
	for _, v := range a {
		result ^= v
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
