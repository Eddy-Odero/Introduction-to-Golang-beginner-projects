package main

import (
	"fmt"
)
func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	dot := -1
	for i, c := range dec {
		if c == '.' {
			if dot != -1 {
				return dec + "\n"
			}
			dot = i
		} else if (c < '0' || c > '9') && !(i == 0 && c == '-') {
			return dec + "\n"
		}
	}

	if dot == -1 {
		return dec + "\n"
	}

	allZero := true
	for i := dot + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			allZero = false
			break
		}
	}
	if allZero {
		return dec + "\n"
	}

	res := dec[:dot] + dec[dot+1:]
	sign := ""
	if res[0] == '-' {
		sign = "-"
		res = res[1:]
	}

	i := 0
	for i < len(res)-1 && res[i] == '0' {
		i++
	}

	return sign + res[i:] + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
