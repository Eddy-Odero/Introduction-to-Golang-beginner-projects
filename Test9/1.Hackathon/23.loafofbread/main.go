package main

import (
	"fmt"
	
)

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	result := ""
	counter := 0
	for i, v := range str {
		if v != ' ' && counter != 5 {
			result += string(v)
			counter++
		} else if counter == 5 {
			result += " "
			counter = 0
		}
		if i == len(str)-1 && len(result) > 0 && result[len(result)-1] == ' ' {
			result = result[:len(result)-1]
		}
	}
	result += "\n"
	return result
}

func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}
