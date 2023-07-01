/*
Given an array A of positive integers. Your task is to find the leaders in the array.
An element of array is leader if it is greater than or equal to all the elements to its right side.
The rightmost element is always a leader.
*/

package main

import (
	"fmt"
)

func main() {
	input := []int{1, 9, 7, 8, 6, 3, 4, 2}

	largest := input[len(input)-1]

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] >= largest {
			largest = input[i]
			fmt.Println(largest)
		}
	}
}
