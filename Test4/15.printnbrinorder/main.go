package main

func PrintNbrInOrder(n int) {
digits := []int{}
if n > 0 {
digits = append(digits, n%10)
n/=10
}
for i:= 0;i < len(digits)-1;i++{
	for j:=0;j < len(digits)-i-1;j++{

	}
}
}
func main(){
PrintNbrInOrder(321)
PrintNbrInOrder(0)
PrintNbrInOrder(321)

}
