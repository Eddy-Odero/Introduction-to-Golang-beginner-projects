package main

func BasicAtoi(s string) int {
	result := 0
	for i:=0;i< len(s);i++{
		if s[i] >='0' && s[i]<= '9'{
			result = result*10 + int(s[i]-'0')
		}else{
			break
		}
	}
return result
}
