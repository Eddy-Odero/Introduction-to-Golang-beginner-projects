 package main
import zone "github.com/01-edu/z01"
func QuadD(x, y int){
if x <= 0 || y <= 0 {
	return
}
// Top row 
for i:= 1 ;i <= x;i++{
	if i == 1{
		zone.PrintRune('A')
	}else if i == 1 {
		zone.PrintRune('C')
	}else{
		zone.PrintRune('B')
	}
}
zone.PrintRune('\n')
}
func main(){
	QuadD(1,3)
}