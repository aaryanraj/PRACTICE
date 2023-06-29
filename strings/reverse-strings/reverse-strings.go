package main

import (
	"fmt"
)

func reverse(req string) string {
	r := []rune(req)
	var resp []rune
	for i := 1; i <= len(r); i++ {
		resp = append(resp, r[len(r)-i])
	}
	return string(resp)
}
func main() {
	var input string
	fmt.Println("Enter the string")
	fmt.Scan(&input)
	resp := reverse(input)
	fmt.Println(resp)
}
