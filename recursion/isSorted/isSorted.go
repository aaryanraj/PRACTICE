package main

import (
	"fmt"
)

func isSorted(arr []int) bool {
	n := len(arr)
	if n == 1 {
		return true
	}

	if arr[n-1] < arr[n-2] {
		return false
	}
	return isSorted(arr[:n-1])
}

func main() {
	var len int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&len)

	var arr []int
	fmt.Println("Enter the elements of the array")
	for i := 0; i < len; i++ {
		var input int
		fmt.Scan(&input)
		arr = append(arr, input)
	}

	fmt.Println(isSorted(arr))
}
