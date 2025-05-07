package main

func StrRev(s string) string {
	runes :=[]rune(s)
k:= len(s)
for i:=0;i<= k/2;i++{
runes[i],runes[k-i-1] =runes[k-i-1],runes[i]	
}
return string(runes)
}
