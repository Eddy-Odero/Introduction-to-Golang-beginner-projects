package main

import (
	"fmt"
	
)

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
n:= len(a)
for i :=0;i < n-1;i++{
	for j:= 0;j< n-i-1;j++{
		if f(a[j], a[j+1]) > 0 {
			a[j],a[j+1] = a[j+1],a[j]
		}
	}
}

}
func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return -1
}
func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)

	fmt.Println(result)
}
