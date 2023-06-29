package main

import (
	"fmt"
	"math"
)

func rev(sl []int) []int {
	fmt.Println("here as well")
	var resp []int
	for i := len(sl) - 1; i >= 0; i-- {
		resp = append(resp, sl[i])
	}
	return resp
}

func main() {
	var answer []int
	input := []int{3, 4, 5, 9, 8, 7, 3, 2}
	k := 3
	n := len(input)
	fmt.Println(n)
	for i := 0; i < n; i += 3 {
		left := i
		right := int(math.Min(float64(i+k), float64(n)))
		rev := rev(input[left:right])
		answer = append(answer, rev...)
	}
	fmt.Println(answer)
}
