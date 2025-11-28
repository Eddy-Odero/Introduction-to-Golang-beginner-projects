package main

import "fmt"

func ToUpper(s string) string {
	var result = " "
	for _, i := range s {
		if i >= 97 && i <= 122 {
			result += string(i - 32)
		} else {
			result+= string(i)
		}

	}

	return result
}
func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
