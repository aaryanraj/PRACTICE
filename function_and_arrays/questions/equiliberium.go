package main

import (
	"fmt"
)

func findEquilibberium(input []int) int {
	var sumLeft, sumRight int

	for i := 1; i < len(input)-1; i++ {
		sumLeft = 0
		sumRight = 0
		for j := i - 1; j >= 0; j-- {
			sumLeft = sumLeft + input[j]
		}

		for k := i + 1; k <= len(input)-1; k++ {
			sumRight = sumRight + input[k]
		}
		if sumLeft == sumRight {
			return input[i]
		}
	}
	return -1
}

func main() {
	input := []int{2, 3, 4, 1, 4, 5}
	ans := findEquilibberium(input)

	fmt.Println(ans)
}
