package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	if s == " "{
		return 0
	}
	result := 0
    sign := 1
    i := 0

    if s[0] == '-' {
        sign = -1
        i++
    }else if s[0] =='+'{
		i++
	}

    for j :=i ; j < len(s);j++ {
        if s[j] < '0' || s[j] > '9' {
			return 0
            
        }
        result = result*10 + int(s[j]-'0')
    }

    return result * sign
}

func add(a, b int) int{
	return (a+b)
}
func sub(a,b int)int{
	return a-b
}
func mult(a, b int)int{
	return a*b
}
func div(a,b int)int{
	
	return a/b
}
func mod(a,b int)int{
	
	return a%b
}

func main(){
	if len(os.Args) < 4 {
		return 
	}
	a:= Atoi(os.Args[1])
	b:= os.Args[2]
	c:=Atoi(os.Args[3])
	

	var op func(int, int) int
	switch b {
	case "+":
		op = add
	case "-":
		op = sub
	case "*":
		op = mult
	case "/":
		op = div
	case "%":
		op = mod
	default:
		return
	}
fmt.Println(op(a,c))
}
