package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	size int
}

func main() {
	var ll LinkedList = LinkedList{}
	for true {
		var operation int
		fmt.Println("Select the operation you want to perform")
		fmt.Println("1. Insertion")
		fmt.Println("2. Deletion")
		fmt.Scan(&operation)
		if operation == 1 {
			var n int
			fmt.Println("How many numbers do you want to enter")
			fmt.Scan(&n)
			for i := 0; i < n; i++ {
				var insert int
				fmt.Println("Select where do you want to insert")
				fmt.Println("1. At the front")
				fmt.Println("2. At the End")
				fmt.Println("3. Somewere in betwen")
				fmt.Scan(&insert)

				var data string
				switch insert {
				case 1:
					{
						fmt.Println("Enter the Data you want to insert")
						fmt.Scan(&data)
						ll.InsertBeginning(data)
					}
				case 2:
					{
						fmt.Println("Enter the Data you want to insert")
						fmt.Scan(&data)
						ll.InsertAtEnd(data)
					}
				case 3:
					{
						var pos int
						fmt.Println("Enter the Data you want to insert")
						fmt.Scan(&data)
						fmt.Println("Enter the position after which you want to insert")
						fmt.Scan(&pos)
						ll.InsertAt(data, pos)
					}

				}
			}
			if err := ll.Display(); err != nil {
				fmt.Printf(err.Error())
			}

		} else if operation == 2 {
			var delete int
			fmt.Println("Select from where do you want to delete")
			fmt.Println("1. From the front")
			fmt.Println("2. From the End")
			fmt.Println("3. From somewere in betwen")
			fmt.Scan(&delete)

			switch delete {
			case 1:
				{
					if val, err := ll.DeleteFirst(); err != nil {
						fmt.Println(err.Error())
						continue
					} else {
						fmt.Println("Deleted value", val, "from beginning of LL")
					}
				}
			case 2:
				{
					if val, err := ll.DeleteLast(); err != nil {
						fmt.Println(err.Error())
						continue
					} else {
						fmt.Println("Deleted value", val, "from the end of LL")
					}
				}
			case 3:
				{
					var pos int
					fmt.Println("Enter the position from where you want to delete")
					fmt.Scan(&pos)
					if val, err := ll.DeleteFromMiddle(pos); err != nil {
						fmt.Println(err.Error())
						continue
					} else {
						fmt.Println("Deleted value", val, "from somewere in the middle of LL")
					}
				}
			}

			if err := ll.Display(); err != nil {
				fmt.Println(err.Error())
			}

		} else {
			fmt.Printf("Invalid input %v, please select from the above option 1 OR 2", operation)
			continue
		}
	}
}

func (dll *LinkedList) CheckIfEmptyAndAdd(node *Node) bool {
	if dll.size == 0 {
		dll.head = node
		dll.size++
		return true
	}
	return false
}

func (dll *LinkedList) InsertBeginning(data interface{}) {
	newNode := &Node{
		data: data,
		prev: nil,
		next: nil,
	}

	if !dll.CheckIfEmptyAndAdd(newNode) {
		head := dll.head
		head.prev = newNode
		newNode.next = head
		dll.head = newNode
		dll.size++
	}
}

func (dll *LinkedList) Display() error {
	if dll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := dll.head
	for current != nil {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

func (dll *LinkedList) InsertAtEnd(data interface{}) {
	newNode := &Node{
		data: data,
		next: nil,
		prev: nil,
	}

	if !dll.CheckIfEmptyAndAdd(newNode) {
		lastNode := dll.head
		for lastNode.next != nil {
			lastNode = lastNode.next
		}

		newNode.prev = lastNode
		lastNode.next = newNode
		dll.size++
	}
}

func (dll *LinkedList) InsertAt(data interface{}, index int) error {
	if index < 1 && index > dll.size+1 {
		return fmt.Errorf("insert: index out of bound")
	}

	if dll.head == nil {
		fmt.Println("InsertAt: List is empty!!!, Creating a new list")
	}

	newNode := &Node{
		data: data,
		next: nil,
		prev: nil,
	}

	if !dll.CheckIfEmptyAndAdd(newNode) {
		current := dll.head

		for i := 1; i < index; i++ {
			current = current.next
		}
		newNode.prev = current.prev
		newNode.next = current

		current.prev.next = newNode
		current.prev = newNode
		dll.size++

	}
	return nil
}

func (dll *LinkedList) DeleteFirst() (interface{}, error) {
	if dll.head == nil || dll.size == 0 {
		return -1, fmt.Errorf("linked list is empty")
	}
	head := dll.head
	if head.prev == nil {
		deleteData := head.data
		dll.head = head.next
		if dll.head != nil {
			dll.head.prev = nil
		}
		dll.size--
		return deleteData, nil
	}
	return -1, fmt.Errorf("something went wrong")
}

// note :- if we have a tail feild in the linked list then deleting last can be brought down to O(1)
func (dll *LinkedList) DeleteLast() (interface{}, error) {
	if dll.head == nil || dll.size == 0 {
		return -1, fmt.Errorf("linked list is empty")
	}

	current := dll.head

	for current.next != nil {
		current = current.next
	}
	if current.prev != nil {
		current.prev.next = nil
	} else {
		dll.head = nil
	}
	current.prev = nil
	dll.size--

	return current.data, nil
}

func (dll *LinkedList) DeleteFromMiddle(pos int) (interface{}, error) {
	if dll.head == nil || dll.size == 0 {
		return -1, fmt.Errorf("list is empty")
	}

	if pos < 1 || pos > dll.size {
		return -1, fmt.Errorf("DeleteFromMiddle: Index out of bound")
	}

	var value interface{}
	var err error
	if pos == 1 {
		if value, err = dll.DeleteFirst(); err != nil {
			return -1, err
		}
		return value, nil
	}

	if pos == dll.size {
		if value, err = dll.DeleteLast(); err != nil {
			return -1, err
		}
		return value, nil
	}

	current := dll.head

	for pos >= 2 && pos != dll.size {
		current = current.next
		pos--
	}

	current.prev.next = current.next
	current.next.prev = current.prev
	dll.size--
	return current.data, nil
}
