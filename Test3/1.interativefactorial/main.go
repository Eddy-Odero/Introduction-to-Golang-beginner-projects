package main
import "fmt"

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 25 {
	}
	if nb == 0 {
		return 1
	}
	factorial := 1
	for i := nb; i >= 2; i-- {
factorial *= i
	}
return factorial
}
func main() {
arg := 4
	fmt.Println(IterativeFactorial(arg))
}