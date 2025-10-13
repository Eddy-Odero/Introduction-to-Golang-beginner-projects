 package main
import zone "github.com/01-edu/z01"

func QuadC(x, y int){
	if x <= 0 || y <= 0 {
		return
	}
	// Top row
	for i := 1; i <= x; i++ {
		if i == 1 {
			zone.PrintRune('A')
		} else if i == x {
			zone.PrintRune('A')
		} else {
			zone.PrintRune('B')
		}
	}
	zone.PrintRune('\n')

	// Middle rows
	for row := 1; row <= y-2; row++ {
		for col := 1; col <= x; col++ {
			if col == 1 || col == x {
				zone.PrintRune('B')
			} else {
				zone.PrintRune(' ')
			}
		}
		zone.PrintRune('\n')
	}
	// Bottom row 
	if y > 1 {
		for i := 1; i <= x; i++ {
			if i == 1 {
				zone.PrintRune('C')
			}else if i == x {
				zone.PrintRune('C')
			} else {
				zone.PrintRune('B')
			}
		}
		zone.PrintRune('\n')
	}
}
	