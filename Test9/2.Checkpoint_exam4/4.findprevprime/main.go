package main

import "fmt"
func IsPrime(nb int) bool {

	if nb < 2 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
for i:= nb; i >= 2;i--{
	if IsPrime(i){
		return i

	}
}
return 0
}
func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
