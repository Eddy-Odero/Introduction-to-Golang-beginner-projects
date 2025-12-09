package main
import "fmt"
func toUpper(r rune)rune{
	if r >= 'a' && r <= 'z' {
		return r-32
	}
	return r
}
func toLower(r rune)rune{
	if r >= 'A' && r <= 'Z' {
		return r +32
	}
	return r
}
func isAlnum(r rune)bool {
	return (r>= 'a' && r <= 'z') ||
	(r >= 'A'&& r <= 'Z') ||
	(r >= '0' && r <= '9')
}
func Capitalize(s string) string {
runes := []rune(s)
inword := false

for i ,r := range runes {
	if isAlnum(r) {
		if !inword {
			runes[i] = toUpper(r)
			inword = true
		
	}else{
		runes[i] = toLower(r)
	}
}else {
	inword = false
}
}
return string(runes)
}

func main(){
fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
