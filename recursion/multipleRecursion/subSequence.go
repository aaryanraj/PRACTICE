/*
array = [1,2,3,4]

Subarray : [1,2],[1,2,3] - is continous and maintains relative order of elements

Subsequence: [1,2,4] - is not continous but maintains relative order of elements

Subset: [1,3,2] - is not continous and does not maintain relative order of elements
*/
// Print all the subsequence of the given array.

package main

import (
	"fmt"
)

func printSubSequence(index int, array, input []int) {
	n := len(input)
	if index == n {
		fmt.Println(array)
		return
	}
	//pick
	array = append(array, input[index])
	printSubSequence(index+1, array, input)
	array = array[:len(array)-1]

	//not pick
	printSubSequence(index+1, array, input)
}

func main() {
	input := []int{3, 1, 2}
	array := []int{}
	printSubSequence(0, array, input)

}

//also refer practice/questions/powerSet.go
