package main
 import (
	"fmt"
	"os"
 )
  
 func isSpace(r byte)bool {
	return r == ' ' || r == '\t'
 }
func main(){
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	s := os.Args[1]
	n := len(s)
	i := 0
	for i < n && isSpace(s[i]) {
		i++
	}
	wordcount := 0
for i < n {
	if wordcount > 0 {
		fmt.Print(" ")
	}
	for i < n && !isSpace(s[i]) {
		fmt.Printf("%c",s[i])
		i++
	}
}
fmt.Println()
}