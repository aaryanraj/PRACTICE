/*package main

import (
	"errors"
	"fmt"
)

type stack struct {
	sl []rune
}

func (s *stack) isEmpty() bool {
	if len(s.sl) == 0 {
		return true
	}
	return false
}

func (s *stack) push(value rune) {
	s.sl = append(s.sl, value)
}

func (s *stack) pop() (rune, error) {
	if s.isEmpty() {
		return 0, errors.New("Stack is empty")
	}
	len := len(s.sl)
	val := s.sl[len-1]
	s.sl = s.sl[:len-1]
	return val, nil
}

func isMatching(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	}
	return false
}
func isOpen(val rune) bool {
	if val == '(' || val == '{' || val == '[' {
		return true
	}

	return false
}
func isClose(val rune) bool {
	if val == ')' || val == '}' || val == ']' {
		return true
	}
	return false
}
func main() {
	var input string
	var st stack
	fmt.Println("Enter the sequence of brackets")
	fmt.Scan(&input)

	inputInRune := []rune(input)
	var balanced bool

	for _, value := range inputInRune {
		if isOpen(value) {
			st.push(value)
			continue
		}
		if !isClose(value) {
			continue
		}
		open, err := st.pop()
		if err != nil {
			balanced = false
		}
		if isMatching(open, value) {
			balanced = true
		} else {
			balanced = false
		}

	}

	if balanced {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Not Balanced")
	}
}
*/