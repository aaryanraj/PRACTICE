package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	var ans []int

	for i, j := 0, len(input)-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			ans = append(ans, input[i])
		} else {
			ans = append(ans, input[j])
			ans = append(ans, input[i])
		}
	}

	fmt.Println(ans)
}
