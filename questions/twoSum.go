/*
Given a list print the index of the numbers whose sum is 13(target)
*/
package main

import (
	"fmt"
)

func main() {
	//input := []int{1, 5, 3, 8, 7}
	sortedInput := []int{1, 3, 4, 8, 9}
	target := 13
	//solution with hashmap; time: O(n), space: O(n)
	/*
		var resp []int
		s := make(map[int]int)
		for index, val := range input {
			if v, found := s[target-val]; found {
				resp = append(resp, v, index)
			}
			s[val] = index

		}
		fmt.Println(resp)
	*/

	//solution with twoPointer; time: O(n), space: O(1)
	//only works on sorted array

	ptr1, ptr2 := 0, len(sortedInput)-1

	for ptr1 < ptr2 {
		current := sortedInput[ptr1] + sortedInput[ptr2]
		if current == target {
			break
		} else if current < target {
			ptr1++
		} else {
			ptr2--
		}
	}

	fmt.Println([]int{ptr1, ptr2})

}
