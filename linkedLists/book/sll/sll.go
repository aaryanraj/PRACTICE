package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (ll *LinkedList) Display() error {
	if ll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

// since we have a size field in the ll and we always update it while adding a element
func (ll *LinkedList) Length() int {
	return ll.size
}

// inserting at the beginning takses O(1) constant time and space
func (ll *LinkedList) InsertBeginning(data interface{}) {
	node := &Node{data: data}

	if ll.head == nil {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.size++
}

// inserting at the end takes O(n) time and O(1) space
func (ll *LinkedList) InsertEnd(data interface{}) {
	node := &Node{data: data,
		next: nil}

	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	ll.size++
}

// inserting a node at Nth position takes O(n) time and O(1) space
func (ll *LinkedList) Insert(data interface{}, pos int) error {
	if pos < 1 || pos > ll.size+1 {
		return fmt.Errorf("insert: index out of bound")
	}
	node := &Node{data: data,
		next: nil}

	var pre, current *Node
	pre = nil
	current = ll.head

	for pos > 1 {
		pre = current
		current = current.next
		pos = pos - 1
	}

	if pre != nil {
		pre.next = node
		node.next = current
	} else {
		node.next = current
		ll.head = node
	}
	ll.size++
	return nil
}

func (ll *LinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFirst: list is empty")
	}
	data := ll.head.data
	ll.head = ll.head.next
	ll.size--
	return data, nil
}

func (ll *LinkedList) DeleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFirst: list is empty")
	}

	var pre *Node
	current := ll.head
	for current.next != nil {
		pre = current
		current = current.next
	}
	if pre != nil {
		pre.next = nil
	} else {
		ll.head = nil
	}
	ll.size--
	return current.data, nil
}

func (ll *LinkedList) DeleteFromMiddle(pos int) (interface{}, error) {
	if pos < 1 || pos > ll.size+1 {
		return fmt.Errorf("deleteFromMiddle: Index out of bound"), nil
	}

	var pre, current *Node
	pre = nil
	position := 0
	current = ll.head
	if pos == 1 {
		ll.head = ll.head.next
	} else {
		for position != pos-1 {
			position = position + 1
			pre = current
			current = current.next
		}
		if current != nil {
			pre.next = current.next
		}
	}
	ll.size--
	return current.data, nil
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
						ll.InsertEnd(data)
					}
				case 3:
					{
						var pos int
						fmt.Println("Enter the Data you want to insert")
						fmt.Scan(&data)
						fmt.Println("Enter the position where you want to insert")
						fmt.Scan(&pos)
						ll.Insert(data, pos)
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
