package piscine

/*
package main

import (

	"fmt"

)

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

	type NodeL struct {
		Data interface{}
		Next *NodeL
	}
*/
func ListAt(l *NodeL, pos int) *NodeL {
	iter := 0
	it := l
	for it != nil {
		if iter == pos {
			return it
		}
		it = it.Next
		iter++
	}
	return nil
}
