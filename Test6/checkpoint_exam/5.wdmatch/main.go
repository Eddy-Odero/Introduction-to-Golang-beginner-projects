package main
import ("os"
)
func main(){
	if len(os.Args) != 3 {
		return 
	}
	
	s1 := []rune(os.Args[1])
	s2 := []rune (os.Args[2])

	i := 0
for _, r := range s2 {
	for i < len(s1) && r == s1[i]{
		i++
	}
	}
	if i == len(s1){

	}
os.Stdout.WriteString(string(s1) + "\n")
}
