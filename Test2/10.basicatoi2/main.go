package main
import "fmt"

func BasicAtoi2(s string) int {
result :=0
for i:=0;i < len(s);i++{
	if s[i] >= '0' && s[i] <='9' {
		result = result*10 + int(s[i]-'0')
	
	}else{
		return 0
	}
}
return result
}
func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}