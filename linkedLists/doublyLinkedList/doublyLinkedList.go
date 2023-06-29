package main

import (
	"fmt"
)

type dllNode struct {
	prev *dllNode
	data int
	next *dllNode
}

type dLinkedList struct {
	head   *dllNode
	length int
}

func (l *dLinkedList) prepend(n *dllNode) {
	newNode := dllNode{
		prev: nil,
		data: n.data,
		next: l.head,
	}
	if l.head != nil {
		l.head.prev = &newNode
	}
	l.head = &newNode
	l.length += 1
}

func (l dLinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length -= 1
	}
	fmt.Printf("\n")
}

func (l dLinkedList) printAdd() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Println(toPrint)
		toPrint = toPrint.next
		l.length -= 1
	}
	fmt.Printf("\n")
}

func main() {
	var len int
	myList := dLinkedList{}
	fmt.Println("enter the length of the linked list")
	fmt.Scan(&len)
	fmt.Println("enter the value to be stored in the list")
	for i := 0; i < len; i++ {
		var value int
		fmt.Scan(&value)
		node := &dllNode{data: value}
		myList.prepend(node)
	}
	myList.printListData()

	myList.printAdd()

}
