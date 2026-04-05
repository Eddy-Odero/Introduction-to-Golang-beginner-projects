package main

import( 
 "github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	count := [10]int{}
	for n > 0 {
		count[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {
		for count[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			count[i]--
		}
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)

}
