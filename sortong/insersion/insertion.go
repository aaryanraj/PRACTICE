//Insertion sort
/*
Algo- take an element and place it in its correct order
*/

package main

import (
	"fmt"
)

func main() {
	input := []int{9, 5, 7, 3, 1, 2, 6}

	for i := 0; i < len(input); i++ {
		for j := i; j > 0 && input[j-1] > input[j]; j-- {
			input[j-1], input[j] = input[j], input[j-1]
		}
	}

	fmt.Println(input)
}
