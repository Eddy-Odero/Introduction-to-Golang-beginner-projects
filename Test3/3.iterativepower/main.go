package main
 import "fmt"

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 1 {
		return nb
	}
	if power == 0 {
		return 1
	}
	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}
func main(){
fmt.Println(IterativePower(4,3))
}