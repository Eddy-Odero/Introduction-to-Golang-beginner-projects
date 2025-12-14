package main
import "fmt"

func isValidBase(base string) bool{
	if len(base) < 2 {
		return false
	}
	for i := 0;i < len(base);i++{
		if base[i] == '+' || base[i]=='-'{
			return false
		}
		for j := i + 1;j < len(base);j++{
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}


func AtoiBase(s string, base string) int {
if !isValidBase(base){
	return 0
}
sign := 1
if len(s) > 0 && s[0] == '-'{
	sign = -1
	s =s[1:]
}
baseLen := len(base)
result := 0
for  _, r := range s {
	index := -1
	for i,b := range base {
		if r == b {
			index= i
			break
		}
	}
	if index == -1 {
		return 0
	}
	result = result*baseLen + index
}
return result * sign
}
func main(){
fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}