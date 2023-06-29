package main

import (
	"fmt"
)

func findDuplicates(arr []int) []int {
	var resp []int
	m := make(map[int]int)

	for _, value := range arr {
		if _, found := m[value]; found {
			m[value]++
			resp = append(resp, value)
		}
		m[value] = 1
	}

	return resp
}

func main() {
	var input int
	fmt.Println("Enter the length of the array")
	fmt.Scan(&input)
	arr := []int{}
	fmt.Println(fmt.Sprintf("Enter %v values", input))
	for i := 0; i < input; i++ {
		var val int
		fmt.Scan(&val)
		arr = append(arr, val)
	}

	resp := findDuplicates(arr)

	fmt.Println(resp)

}
