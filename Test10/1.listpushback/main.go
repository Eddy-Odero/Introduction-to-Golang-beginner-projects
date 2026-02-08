package main
import "fmt"

type nodeL struct{
	Data interface{}
	Next *nodeL
}
type List struct{
	Head *nodeL
	Tail *nodeL
}
func ListPushBack(l *List, data interface{}) {
n := &nodeL{Data:data}
 
if l.Head == nil {
	l.Head = n
	l.Tail = n
	return 

}
l.Tail.Next = n
l.Tail = n
}
func main() {

	link := &List{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
 	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}