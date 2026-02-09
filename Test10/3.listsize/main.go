package main
import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
if l.Head == nil{
	l.Head = n
	l.Tail = n
	return
}
n.Next = l.Head
l.Head = n
}

func ListSize(l *List) int {
count := 0
n := l.Head
for n != nil{
count++
n = n.Next
}
return  count
}

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}