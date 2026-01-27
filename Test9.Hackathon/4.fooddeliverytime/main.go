package main

import (
	"fmt"

)
type food struct {
  preptime int
}
var list = map[string]food{
	"burger" : {preptime:15},
	"chips" : {preptime:10},
	"nuggets" : {preptime:12},
}
func FoodDeliveryTime(order string) int {
	 i,j := list[order]
		if !j{
return -1
		}
		return i.preptime
	}


func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}
