package main

import (
	"fmt"
)
func Join(strs []string, sep string) string {
	if len(strs)== 0{
		return ""
	}
result :=strs[0]
for i := 0;i< len(strs);i++{
	result+=sep +strs[i]
}
return result
}
func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
