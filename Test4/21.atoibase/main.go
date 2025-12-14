package main
func AtoiBase(s string, base string) int {
if !isvalidBase(base){
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
	if index == -1{
		return 0
	}
	result = result*baseLen + index
}
return result * sign
}
func main(){

}