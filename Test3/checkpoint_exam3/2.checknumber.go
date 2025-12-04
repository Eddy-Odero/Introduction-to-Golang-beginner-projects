package main

func CheckNumber(arg string) bool {
	for _, i := range arg {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {

		} else if i >= '0' && i <= '9' {
			return true
		}
	}
	return false
}
