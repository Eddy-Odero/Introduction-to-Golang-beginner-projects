package main
import "fmt"

func BasicJoin(elems []string)string{
	result := ""
	for _, s :=range elems{
		result +=s
	}
	return result
}
func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}