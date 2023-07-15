//find the k th larhest element in an array

package main

import (
	"fmt"
	"sort"
)

/*
func main() {
	input := []int{9, 5, 7, 3, 4, 6}
	k := 3
	kth := 0
	sort.Sort(sort.Reverse(sort.IntSlice(input)))
	fmt.Println(input)
	for i := 0; i <= k-1; i++ {
		kth = input[i]
	}
	fmt.Println(kth)
}
*/

func main() {
	input := []int{9, 9, 9, 9, 4, 6} //{2,4,6,7,8,9}
	sort.Ints(input)
	k := 3
	var count int = 1
	var found bool
	mapOfInput := make(map[int]int)
	for i := len(input) - 1; i >= 0; i-- {
		if _, ok := mapOfInput[input[i]]; !ok {
			mapOfInput[input[i]] = count
			count++
		}
	}

	for key, value := range mapOfInput {
		if value == k {
			fmt.Println(key)
			found = true
		}
	}
	fmt.Println(mapOfInput)
	if !found {
		fmt.Println("-1")
	}

}
