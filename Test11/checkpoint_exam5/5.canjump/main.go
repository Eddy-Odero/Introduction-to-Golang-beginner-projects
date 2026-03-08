package main

import (
	"fmt"
	
)
func CanJump(a []uint) bool {
	if len(a) == 0 {
		return false
	}

	i := 0
	last := len(a) - 1

	for i < len(a) {
		if i == last {
			return true
		}
		if a[i] == 0 || i+int(a[i]) > last {
			return false
		}
		i += int(a[i])
	}
	return false
}
func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
