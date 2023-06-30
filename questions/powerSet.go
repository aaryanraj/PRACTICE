// find all the combinaltions possible for a given string. continous and non-continous (subSequence)
package main

import (
	"fmt"
	"math"
)

func main() {
	input := "abcd"
	n := len(input)
	var count int = 1
	for num := 0; num < int(math.Pow(2, float64(n))); num++ {
		var str string
		for i := 0; i < n; i++ {
			if num&(1<<i) == 0 {
				str = str + string(input[i])
			}
		}
		fmt.Println(count, str)
		count++
	}
}
