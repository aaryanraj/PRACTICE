package main

import "fmt"

func maxHeapify(a []int, n int) []int {
	l := 2*n + 1
	r := 2*n + 2
	var largest = n
	if l <= len(a) && a[l] > a[n] {
		largest = l
	}
	if r <= len(a) && a[r] > a[n] && a[r] > a[l] {
		largest = r
	}
	if largest != n {
		a[n], a[largest] = a[largest], a[n]
		a = maxHeapify(a, largest)
	}
	return a
}

func main() {
	input := []int{28, 30, 20, 15, 10, 8, 16}
	ans := maxHeapify(input, 0)
	fmt.Println(ans)
}
