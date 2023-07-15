package main

import (
	"fmt"
)

func main() {
	input := []int{9, 7, 2, 5, 4, 3}
	target := 10

	s := make(map[int]int)
	var resp []int
	for index, value := range input {
		if v, found := s[target-value]; found {
			resp = append(resp, v, index)
		} else {
			s[value] = index
		}
	}
	fmt.Println(resp)
}
