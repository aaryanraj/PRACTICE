//Selection sort
/*
Algo- itrate over the array, find the minimum element and swap with i, redo with next index
*/

package main

import (
	"fmt"
)

func main() {
	input := []int{9, 5, 7, 1, 6, 8, 2}

	for i := 0; i < len(input)-1; i++ {
		min := i
		for j := i; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}

		temp := input[i]
		input[i] = input[min]
		input[min] = temp
	}

	fmt.Println(input)
}
