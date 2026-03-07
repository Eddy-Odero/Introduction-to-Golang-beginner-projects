package main

import (
	"fmt"

)
func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	step := 1
	if from > to {
		step = -1
	}

	for i := from; ; i += step {
		result += string('0'+(i/10)) + string('0'+(i%10))

		if i == to {
			break
		}

		result += ", "
	}

	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
