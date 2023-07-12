//Merge sort
/*
Algo- divide the given array in halfs using recursion until it is possible, the merge in sorted order.
*/

package main

import (
	"fmt"
)

func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	ans := []int{}
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			ans = append(ans, arr1[i])
			i++
		} else {
			ans = append(ans, arr2[j])
			j++
		}
	}

	if i < len(arr1) {
		ans = append(ans, arr1[i:]...)
	}

	if j < len(arr2) {
		ans = append(ans, arr2[j:]...)
	}
	return ans
}

func mergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	first := mergeSort(input[:len(input)/2])
	second := mergeSort(input[len(input)/2:])
	return merge(first, second)
}

func main() {
	input := []int{9, 2, 8, 4, 3, 6, 1, 1, 2}
	ans := mergeSort(input)
	fmt.Println(ans)
}
