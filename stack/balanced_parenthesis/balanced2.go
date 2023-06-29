package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

func (s *stack) push(val rune) {
	s.sl = append(s.sl, val)
}

func (s *stack) pop() (val rune, err error) {
	var resp rune
	if s.isEmpty() {
		return resp, errors.New("Stack is empty")
	}
	top := s.sl[len(s.sl)-1]
	s.sl = s.sl[:len(s.sl)-1]
	return top, nil
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the text to be checked")
	scanner.Scan()
	input := scanner.Text()
	runeOfInput := []rune(input)
	var balanced bool
	var st stack
	for i := 0; i <= len(runeOfInput)-1; i++ {
		if runeOfInput[i] == '(' || runeOfInput[i] == '{' || runeOfInput[i] == '[' {
			st.push(runeOfInput[i])
			continue
		}

		if runeOfInput[i] != ')' && runeOfInput[i] != '}' && runeOfInput[i] != ']' {
			continue
		}

		val, err := st.pop()
		if err != nil {
			balanced = false
		}

		balanced = isMatching(val, runeOfInput[i])
	}

	if balanced {
		fmt.Println("The String is Perfectly balanced")
	} else {
		fmt.Println("The String is not balanced")
	}

}
