package main

import (
	"fmt"

)
func Compact(ptr *[]string) int {
var result []string
for _, i := range *ptr {
	if i != ""{
result = append(result, i)

	}

}
*ptr = result
return len(result)
}
const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
