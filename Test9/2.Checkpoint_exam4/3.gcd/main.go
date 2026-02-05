package main

import (
	"fmt"
	
)
func Gcd(a, b uint) uint {
a= a%b
if b% 2 == 0{

}
return b
}
func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
