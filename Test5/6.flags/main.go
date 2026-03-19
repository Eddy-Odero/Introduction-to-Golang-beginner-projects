package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printHelp() {
	printStr("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n")
	printStr("--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}
	var insertStr string
	var mainStr string
	var orderFlag bool

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		} else if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--insert" || arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		} else {
			mainStr = arg
		}
	}
	combined := mainStr + insertStr
	res := []rune(combined)
	if orderFlag {
		for i := 0; i < len(res)-1; i++ {
			for j := 0; j < len(res)-i-1; j++ {
				if res[j] > res[j+1] {
					res[j], res[j+1] = res[j+1], res[j]
				}
			}
		}
	}
	for _, r := range res {
		z01.PrintRune(r)
	}
	if len(res) > 0 {
		z01.PrintRune('\n')
	}
}
