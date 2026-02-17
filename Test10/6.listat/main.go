package main

import (
	"fmt"
	
)
type NodeL struct{
	Data interface{}
	Next *NodeL
}
type List struct{
	Head *NodeL
	Tail *NodeL
}
func ListPushBack(l *List, data interface{}) {
n := &NodeL{Data:data}
 
if l.Head == nil {
	l.Head = n
	l.Tail = n
	return 

}
l.Tail.Next = n
l.Tail = n
}

func ListAt(n *NodeL, pos int) *NodeL {
	if n == nil || pos < 0 {
		return nil
	}

	i := 0
	current := n

	for current != nil {
		if i == pos {
			return current
		}
		current = current.Next
		i++
	}

	return nil
}

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
}
