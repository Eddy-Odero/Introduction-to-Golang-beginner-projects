package main

import (
	"fmt"
	"os"
)
 func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
 }
func atoi(s string) (int,bool){
	runes := []rune(s)
	if len(runes) == 0 {
    return 0, false
	}
result := 0
for _, r := range runes{
	if  !isDigit(r){
		return 0,false
	}
	result = result*10 +int(r-'0')
}
   return result,true
}
func isPrime(n int)bool{
	if n < 2 {
		return false
	}
	for i := 2 ; i*i <= n ; i++{
		if n%i== 0{
		return false
		}
	}
	return true
}
func main(){
	if len(os.Args) != 2{
		fmt.Println(0)
		return
	}
	n, ok := atoi(os.Args[1])
	if !ok || n <= 0 {
fmt.Print(0)
return
	}
	sum := 0
	for i:= 2 ; i <= n;i++{
		if isPrime(i) {
           sum +=i
		}
		
	}
	fmt.Println(sum)
}