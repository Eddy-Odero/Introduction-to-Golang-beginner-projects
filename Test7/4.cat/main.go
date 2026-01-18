package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		data, err := io.ReadAll(os.Stdin)
		check(err)
		os.Stdout.Write(data)
		return
	}
	for _, filename := range args {
		data, err := os.ReadFile(filename)
		check(err)
		os.Stdout.Write(data)
	}
}

func check(err error) {
	if err != nil {
		os.Stdout.WriteString("ERROR: ")
		os.Stdout.WriteString(err.Error())
		os.Stdout.WriteString("\n")
		os.Exit(1)
	}
}
