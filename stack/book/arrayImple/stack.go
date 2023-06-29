package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top      int
	capacity uint
	array    []interface{}
}

func (stack *Stack) Init(capacity uint) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]interface{}, capacity)
	return stack
}

func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

func (stack *Stack) Size() uint {
	return uint(stack.top) + 1
}

func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return errors.New("Stack is full")
	}
	stack.top++
	stack.array[stack.top] = data
	return nil
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is Empty")
	}
	temp := stack.array[stack.top]
	stack.top--
	return temp, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is Empty")
	}
	return stack.array[stack.top], nil
}

func (stack *Stack) Drain() {
	stack.array = nil
	stack.top = -1
}

func main() {
	var size uint
	fmt.Println("Enter the size of the Stack")
	fmt.Scan(&size)

	stack := NewStack(size)
	var op int
	for true {
		fmt.Println("Select the operation you want to perform")
		fmt.Println("1. PUSH")
		fmt.Println("2. POP")
		fmt.Scan(&op)
	}

	for i := 0; i < int(size); i++ {
		var value string
		fmt.Println("Enter the Value to be pushed")
		fmt.Scan(&value)
		stack.Push(value)
	}
	var data interface{}
	var err error
	if data, err = stack.Peek(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

}
