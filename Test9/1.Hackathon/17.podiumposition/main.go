package main

import (
    "fmt"
   
)
func PodiumPosition(podium [][]string) [][]string {
n :=len(podium)

for i:=0;i < n/2;i++{
podium[i],podium[n-i-1] =podium[n-i-1],podium[i]
}
return podium
}



func main() {

    p4 := []string{"4th Place"}
    p3 := []string{"3rd Place"}
    p2 := []string{"2nd Place"}
    p1 := []string{"1st Place"}

    position := [][]string{p4, p3, p2, p1}
    fmt.Println(PodiumPosition(position))
}


