package main 

func DigitLen(n, base int) int {
	count :=0
	if base < 2 || base > 36 {
		return -1
	}
	if n  < 0{ 
		n = -n
	}

for n > 0{
	n/=base
	count++
}
return count
}
