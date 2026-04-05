package main

func LastRune(s string) rune {
	n := []rune(s)
	return n[len(s)-1]
}
