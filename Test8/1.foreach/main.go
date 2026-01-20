package main
import "github.com/01-edu/z01"

func printNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func ForEach(f func(int), a []int){ 
for _, i := range a {
f(i)
}

}

func main(){
a:= []int{1, 2, 3, 4, 5, 6}
ForEach(printNbr, a)
}
