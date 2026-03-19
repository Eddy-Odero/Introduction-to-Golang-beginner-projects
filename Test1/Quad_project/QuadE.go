package main

import (
	"github.com/01-edu/z01"
	
)

func QuadE(x,y int){
	if x == 0 || y == 0 {
		return 
	}
	//top
	for i := 1; i <= x; i ++{
		if i == 1{
			z01.PrintRune('A')
		}else if i == x{
			z01.PrintRune('C')
		}else{
			z01.PrintRune('B')
		}
		
	}
	z01.PrintRune('\n')
	//middle
	for row:= 1; row <= y-2; row++{
		for col := 1;col <= x;col++{
			if col == 1 || col == x{
				z01.PrintRune('B')
			}else{
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
	// bottom
	if y > 1 {
		 for i := 1; i <= x;i++ {
			if i == 1 {
				z01.PrintRune('C')
			}else if i == x{
				z01.PrintRune('A')
			}else{
				z01.PrintRune('B')
			}
		}
		z01.PrintRune('\n')
	}
	
}
func main(){
	QuadE(5,3)
}