// find the secondLargestelement in the list and print the number as well as the index.
// if the number is duplicate print the index of the first occurance.
package main

import (
	"fmt"
)

func main() {
	input := []int{5, 9, 7, 8, 6, 4, 8, 10}
	var first, second int = input[0], input[0]
	var index int
	for i := 0; i < len(input); i++ {
		if input[i] > first {
			second = first
			first = input[i]

		} else if input[i] >= second && input[i] < first {
			second = input[i]
		}
	}

	for j := 0; j < len(input)-1; j++ {
		if input[j] == second {
			index = j
			break
		}
	}
	fmt.Println(first, second, index)
}
