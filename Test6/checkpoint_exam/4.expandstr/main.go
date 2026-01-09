package main

import "os"

func main() {
	if len(os.Args) < 2 {
		return
	}
	inword := false 
	count := 0
	for _, arg := range os.Args[1:]{
		for i := 0 ; i < len(arg) ;i++{
c := arg[i]
if c > 32 && c < 126 {
	if !inword{
		if count > 0 {
			os.Stdout.Write([]byte{' '})
		}
		inword = true 
		count++
	}
	os.Stdout.Write([]byte{c})
}else{
	inword = false
}
		}
		inword = false
	}
}
