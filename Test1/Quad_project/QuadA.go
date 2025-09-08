package main

import zone "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	// Top row
	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			zone.PrintRune('o') // Corners
		} else {
			zone.PrintRune('-') // Top edge
		}
	}
	zone.PrintRune('\n')

	// Middle rows (if any)
	for row := 1; row <= y-2; row++ {
		for col := 1; col <= x; col++ {
			if col == 1 || col == x {
				zone.PrintRune('|') // Left and right edges
			} else {
				zone.PrintRune(' ') // Inside
			}
		}
		zone.PrintRune('\n')
	}

	// Bottom row (only if y > 1)
	if y > 1 {
		for i := 1; i <= x; i++ {
			if i == 1 || i == x {
				zone.PrintRune('o') // Corners
			} else {
				zone.PrintRune('-') // Bottom edge
			}
		}
		zone.PrintRune('\n')
	}
}


