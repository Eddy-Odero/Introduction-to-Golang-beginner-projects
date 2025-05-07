package main

func Atoi(s string) int {
	if s == " "{
		return 0
	}
	result := 0
    sign := 1
    i := 0

    if s[0] == '-' {
        sign = -1
        i++
    }else if s[0] =='+'{
		i++
	}

    for j :=i ; j < len(s);j++ {
        if s[j] < '0' || s[j] > '9' {
			return 0
            
        }
        result = result*10 + int(s[j]-'0')
    }

    return result * sign
}




