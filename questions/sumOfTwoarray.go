package main

import (
	"fmt"
)

func arrToNum(input []int) int {
	num, s := 0, 1
	for i := len(input) - 1; i >= 0; i-- {
		num = num + input[i]*s
		s = s * 10
	}
	return num
}
func main() {
	a := []int{3, 4, 5}
	b := []int{5, 8, 9}

	aToInt := arrToNum(a)
	bToInt := arrToNum(b)

	fmt.Println(aToInt + bToInt)

}
