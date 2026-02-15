package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	inputBytes, _ := io.ReadAll(os.Stdin)

	// If no input at all → print nothing
	if len(inputBytes) == 0 {
		return
	}

	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	// Remove last empty line caused by trailing newline
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	height := len(lines)
	if height == 0 {
		fmt.Println("Not a quad function")
		return
	}

	width := len(lines[0])

	// Ensure all lines have same width
	for _, line := range lines {
		if len(line) != width {
			fmt.Println("Not a quad function")
			return
		}
	}

	var matches []string

	if quadA(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", width, height))
	}
	if quadB(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", width, height))
	}
	if quadC(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", width, height))
	}
	if quadD(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", width, height))
	}
	if quadE(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", width, height))
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	sort.Strings(matches)
	fmt.Println(strings.Join(matches, " || "))
}

func buildQuad(width, height int, tl, tr, bl, br, h, v rune) string {
	result := ""

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch {
			case y == 0 && x == 0:
				result += string(tl)
			case y == 0 && x == width-1:
				result += string(tr)
			case y == height-1 && x == 0:
				result += string(bl)
			case y == height-1 && x == width-1:
				result += string(br)
			case y == 0 || y == height-1:
				result += string(h)
			case x == 0 || x == width-1:
				result += string(v)
			default:
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadA(w, h int) string {
	return buildQuad(w, h, 'o', 'o', 'o', 'o', '-', '|')
}

func quadB(w, h int) string {
	return buildQuad(w, h, '/', '\\', '\\', '/', '*', '*')
}

func quadC(w, h int) string {
	return buildQuad(w, h, 'A', 'A', 'C', 'C', 'B', 'B')
}

func quadD(w, h int) string {
	return buildQuad(w, h, 'A', 'C', 'A', 'C', 'B', 'B')
}

func quadE(w, h int) string {
	return buildQuad(w, h, 'A', 'C', 'C', 'A', 'B', 'B')
}

