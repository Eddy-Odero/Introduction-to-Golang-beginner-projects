package main

import "fmt"

func PrintNbrInOrder(n int) {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	for i := 0; i < len(digits)-1; i++ {
		for j := 0; j < len(digits)-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
	result := 0
	for _, d := range digits {
		result = result*10 + d
	}
	fmt.Print(result)
}
func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)

}
