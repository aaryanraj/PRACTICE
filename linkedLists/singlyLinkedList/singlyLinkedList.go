package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length += 1
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length -= 1
	}
	fmt.Printf("\n")
}

func (l linkedList) printAdd() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Println(toPrint)
		toPrint = toPrint.next
		l.length -= 1
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	//edge cases
	//if the linkedList is empty
	if l.length == 0 {
		return
	}

	//if the current head has the value to be removed
	if l.head.data == value {
		l.head = l.head.next
		l.length -= 1
		return
	}

	//for a singly linked list since it only has the info of the next node so we will use the previous
	//node to find the value that needs to be deleted
	preToDelete := l.head
	for preToDelete.next.data != value {
		if preToDelete.next.next != nil {
			preToDelete = preToDelete.next
		} else {
			fmt.Println("value not found in the list")
			return
		}
	}
	preToDelete.next = preToDelete.next.next
	l.length -= 1
}

func ReverseList(head *node) {
	if head == nil {
		return
	}
	ReverseList(head.next)
	fmt.Printf("%d ", head.data)

}
func main() {
	var len int
	myList := linkedList{}
	fmt.Println("enter the length of the linked list")
	fmt.Scan(&len)
	fmt.Println("enter the value to be stored in the list")
	for i := 0; i < len; i++ {
		var value int
		fmt.Scan(&value)
		node := &node{data: value}
		myList.prepend(node)
	}
	myList.printListData()

	var del int
	fmt.Println("Enter the value you want to remove from list")
	fmt.Scan(&del)
	myList.deleteWithValue(del)
	myList.printListData()
	fmt.Println("Reversed linked list is")
	ReverseList(myList.head)
	fmt.Printf("\n")
	fmt.Println("Note: here the list is only printed in reverse order, the underlying list is untouched")
	myList.printListData()
	myList.printAdd()
}
