package main

import (
	"fmt"
	
)
func SortWordArr(a []string) {
	i := len(a)
	for j:=0;j < i-1;j++{
		for k :=0;k < i-j-1;k++{
			if a[k] > a[k+1]{
				a[k],a[k+1] = a[k+1], a[k]
			}
		}
	}

}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
fmt.Println(result)
}