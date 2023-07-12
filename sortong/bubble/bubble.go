//Bubble sort
/*
pushes the maximun to the last by adjacent swapping , opposite to selection sort.
*/

package main

import (
	"fmt"
)

func main() {
	input := []int{8, 9, 4, 3, 1, 5, 2}
	for k := 1; k < len(input); k++ { // in every itration of k we shift max at last
		didSwap := 0
		for i := 0; i < len(input)-k; i++ { // in every itration of i, we swap if input[i]<input[i+1]
			if input[i] > input[i+1] {
				input[i+1], input[i] = input[i], input[i+1]
				didSwap++
			}
		}
		if didSwap == 0 {
			break
		}
	}
	fmt.Println(input)
}
