package main

import (
	"fmt"
)
func Abort(a, b, c, d, e int) int {
	arr:= []int{a,b,c,d,e}
	n := len(arr)
	for i := 0;i <n-1;i++{
		for j:=0;j< n-i-1;j++{
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
mid := n/2
if n%2 !=0{
	
}
return (arr[mid])
}
func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
