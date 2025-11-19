package main

import "fmt"


func IsUpper(s string) bool {
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("Heh!"))

}
