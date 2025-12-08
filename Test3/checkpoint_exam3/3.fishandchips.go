package main



func FishAndChips(n int) string {
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	} else if n%2 == 0 {
		return "fish"
	}
	return "chips"
}

