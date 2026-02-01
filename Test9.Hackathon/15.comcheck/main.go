package main

import (
	"fmt"
	"os"
)
func main(){
	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1:]
	for _, ch := range args {
		if ch == "01"  || ch == "galaxy" || ch == "galaxy 01"{
		fmt.Println("Alert!!!")
		return
	}
}
}