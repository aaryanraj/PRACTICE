/*
You are given an array A, of N elements. Find minimum index based distance between two elements of the array,
*/
package main

import (
	"errors"
	"fmt"
)

func minDistance(in []int, x, y int) (int, error) {
	start, end := -1, -1
	dist := 0
	for i := 0; i < len(in); i++ {
		if in[i] == x {
			start = i
		}

		if in[i] == y {
			end = i
		}

		if start != -1 && end != -1 {
			dist = end - start
			return dist, nil
		}
	}
	return 0, errors.New("Invald input")
}

func main() {
	input := []int{1, 2, 5, 4, 6, 2}
	x, y := 1, 6
	ans, error := minDistance(input, x, y)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(ans)
	}
}
