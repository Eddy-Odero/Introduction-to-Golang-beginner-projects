package main

import (
	

	"github.com/01-edu/z01"
)
func Rot14(s string) string {
result := []rune {}
for _, i := range s{
	if i >= 'a' && i <='z'{
		shift := ((i-'a' +14)%26) + 'a'
		result = append(result, shift)

	}else if i >= 'A' && i <= 'Z'{
		shift := ((i-'A' + 14 )%26)+'A'
		result = append(result, shift)
	}else{
		result = append(result, i)
	}
}
return string(result)
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
