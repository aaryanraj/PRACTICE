package main

import "fmt"

func bubble(arr []int, rng int) []int {
	if rng == 1 {
		return arr
	}
	for i := 0; i < rng; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	return bubble(arr, rng-1)
}
func main() {
	input := []int{9, 5, 8, 2, 3, 4}
	ans := bubble(input, len(input)-1)
	fmt.Println(ans)
}
