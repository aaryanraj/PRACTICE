package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{1, 5, 3, 2, 4, 7}

	sort.Ints(input)
	var missing int
	for i := 0; i < len(input)-1; i++ {
		if input[i]+1 != input[i+1] {
			missing = i + 2
		}
	}

	fmt.Println(missing)
}
