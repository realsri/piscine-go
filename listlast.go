package piscine

//package main
/*
import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{data, nil}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}
*/
func ListLast(l *List) interface{} {
	/*last := l.Head
	for last != nil {
		last = last.Next
	}
	return last*/
	if l.Tail != nil {
		return l.Tail.Data
	} else {
		return nil
	}
}

/*
func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}
*/
