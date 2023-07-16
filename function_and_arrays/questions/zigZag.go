package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{8, 3, 9, 4, 7, 5, 1}

	sort.Ints(input)

	for i := 1; i < len(input)-1; i += 2 {
		input[i], input[i+1] = input[i+1], input[i]
	}

	fmt.Println(input)
}
