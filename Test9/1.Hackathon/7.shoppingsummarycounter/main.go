package main

import (
	"fmt"
	
)
func ShoppingSummaryCounter(str string) map[string]int {
	count := make(map[string]int)
	word := ""
	for _,ch := range str{
		if ch == ' '{
		if word != ""{
			count[word]++
			word = ""
		}
	}else{
		word+= string(ch)
	}

}
if word != "" {
	count[word]++
}
return count
}
func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
