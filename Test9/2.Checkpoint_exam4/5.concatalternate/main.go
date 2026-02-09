package main 
import "fmt"

func ConcatAlternate(slice1,slice2 []int) []int {
	var first,second  []int

	if len(slice2) > len(slice1){
		first = slice2
		second = slice1
	}else{
	first = slice1
	second = slice2

}
var result []int
for i := 0; i < len(first) || i < len(second); i++{
	if i < len(first){
		result = append(result, first[i])
	}
	if i < len(second){
		result = append(result, second[i])
	}
}
return result
}
func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
