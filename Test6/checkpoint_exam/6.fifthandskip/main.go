package main

import (
	"fmt"
	
)
func FifthAndSkip(str string) string {
if len(str)== 0 {

return "\n"
}
if len(str) < 3 {
	return "invalid input"
}
var result string
var current string
 for _, ch := range str {
	if ch != ' '{
current +=string(ch)
	}
	if len(current) == 5 {
		result +=(current + " ")
		current = " "
	}
 }
 if len(result) > 0 {
	result = result[:len(result)-1]
 }
 return result
}
func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
