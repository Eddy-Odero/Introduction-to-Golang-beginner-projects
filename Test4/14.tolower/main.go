package main

import "fmt"

func ToLower(s string) string {
	var result = " "

	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			result += string(i + 32)
		} else {
			result += string(i)
		
		}
	}
	return result
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))

}
