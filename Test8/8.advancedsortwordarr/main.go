package main

import (
	"fmt"
	
)

func AdvancedSortWordArr(a []string, f func(a, b string) int) {


}
func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a > b {
		return -1
	}
	return 1
}
func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)

	fmt.Println(result)
}
