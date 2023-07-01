//Given an unsorted array arr[] of size N.
//Rotate the array to the left (counter-clockwise direction) by D steps, where D is a positive integer.

package main

import (
	"fmt"
)

func rotate(in []int, d, n int) []int {
	for i := 0; i < d; i++ {
		temp := in[0]
		in = in[1:]
		in = append(in, temp)
	}
	return in
}

func main() {
	input := []int{2, 5, 1, 3, 8}
	n := len(input)
	d := 2

	ans := rotate(input, d, n)
	fmt.Println(ans)
}
