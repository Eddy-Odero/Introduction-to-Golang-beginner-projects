package main

import "github.com/01-edu/z01"

func QuadB(x, y int) {
    if x <= 0 || y <= 0 {
        return
    }

    // Top row
    for i := 1; i <= x; i++ {
        if i == 1 {
            z01.PrintRune('/')
        } else if i == x {
            z01.PrintRune('\\')
        } else {
            z01.PrintRune('*')
        }
    }
    z01.PrintRune('\n')

    // Middle rows
    for row := 1; row <= y-2; row++ {
        for col := 1; col <= x; col++ {
            if col == 1 || col == x {
                z01.PrintRune('*')
            } else {
                z01.PrintRune(' ')
            }
        }
        z01.PrintRune('\n')
    }

    // Bottom row (only if y > 1)
    if y > 1 {
        for i := 1; i <= x; i++ {
            if i == 1 {
                z01.PrintRune('\\')
            } else if i == x {
                z01.PrintRune('/')
            } else {
                z01.PrintRune('*')
            }
        }
        z01.PrintRune('\n')
    }
}

func main() {
    QuadB(1, 5)
}