package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Usage: go run . -c <numBytes> <file1> [<file2> ...]")
		os.Exit(1)
	}
	if args[0] != "-c" {
		fmt.Println("Error: First argument must be '-c'")
		os.Exit(1)
	}
	numBytes := 0
	for _, char := range args[1] {
		if char < '0' || char > '9' {
			fmt.Println("Error: Second argument must be a positive integer")
			os.Exit(1)
		}
		numBytes = numBytes*10 + int(char-'0')
	}
	exitCode := 0

	for _, fileName := range args[2:] {
		fmt.Printf("\n==> %s <==\n", fileName)
		if err := ztail(fileName, numBytes); err != nil {
			fmt.Println(err)
			exitCode = 1
		}
	}
	os.Exit(exitCode)
}

func ztail(fileName string, numBytes int) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Error opening file %s: %v", fileName, err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Error getting file stats %s: %v", fileName, err)
	}
	fileSize := stat.Size()
	if fileSize < int64(numBytes) {
		numBytes = int(fileSize)
	}
	buffer := make([]byte, numBytes)
	_, err = file.ReadAt(buffer, fileSize-int64(numBytes))
	if err != nil {
		return fmt.Errorf("Error reading file %s: %v", fileName, err)
	}
	fmt.Print(string(buffer))
	return nil
}

