package main

import (
	"fmt"
)
func DescendAppendRange(max, min int) []int {
	var result []int
	if min >= max {
		return result
	}
	for i := max ; i > min;i--{
result = append(result,i )
	}
return result
}

func main() {
	fmt.Println(DescendAppendRange(10, 5))
	fmt.Println(DescendAppendRange(5, 10))
}
