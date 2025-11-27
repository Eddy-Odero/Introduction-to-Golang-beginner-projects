package main

import "fmt"

func IsPrintable(s string) bool {
	for _, i := range s {
		if i >= 32 && i <= 126{

		} else {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}
