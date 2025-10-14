package main 
import zone "github.com/01-edu/z01"

func QuadE(x,y int){
	if x == 0 || y == 0 {
		return 
	}
	for i := 1; i <= x; i ++{
		if i == 1{
			zone.PrintRune('A')
		}else if i == x{
			zone.PrintRune('C')
		}else{
			zone.PrintRune('B')
		}
	}
}
func main(){
	QuadE(5,1)
}