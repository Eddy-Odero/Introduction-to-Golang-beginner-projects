package main
import "fmt"
func IsLower(s string) bool {
for _,i := range s{
	if i >= 'a' && i <='z'{

	}else{
		return false
	}
}
return true
}
func main(){
fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}
