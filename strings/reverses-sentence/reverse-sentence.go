package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(words []string) []string {
	var resp []string
	for i := 1; i <= len(words); i++ {
		resp = append(resp, words[len(words)-i])
	}
	return resp
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	fmt.Println("Enter the sentence to be reversed")
	scanner.Scan()
	input = scanner.Text()

	words := strings.Split(input, " ")
	resp := reverse(words)
	revStr := strings.Join(resp, " ")

	fmt.Println(revStr)
}
