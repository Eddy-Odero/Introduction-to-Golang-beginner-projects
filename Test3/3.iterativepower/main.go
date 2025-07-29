package main
 import "fmt"
 func Iterativepower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb *Iterativepower(nb,power-1)
 }
func main(){
fmt.Println(Iterativepower(4,3))
}