package main

func PrintNbrInOrder(n int) {
digits := []int{}
if n > 0 {
digits= append(digits, n%10)
n/=10
}
}
func main(){
PrintNbrInOrder(321)
PrintNbrInOrder(0)
PrintNbrInOrder(321)

}
