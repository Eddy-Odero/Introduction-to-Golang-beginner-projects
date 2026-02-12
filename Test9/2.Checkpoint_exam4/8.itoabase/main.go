package main
func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}

	digits := "0123456789ABCDEF"
	negative := false

	if value < 0 {
		negative = true
		value = -value
	}

	result := ""

	for value > 0 {
		remainder := value % base
		result = string(digits[remainder]) + result
		value /= base
	}

	if negative {
		result = "-" + result
	}

	return result
}

func main(){
	ItoaBase(42, 2) 
ItoaBase(42, 16)  
ItoaBase(-42, 16) 
ItoaBase(255, 16) 

}