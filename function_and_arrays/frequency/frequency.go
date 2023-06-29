package main

import (
	"fmt"
)

func main() {
	var number int64
	var input int
	fmt.Println("Enter the number")
	fmt.Scan(&number)
	fmt.Println("Enter the digit")
	fmt.Scan(&input)

	freq := getFrequency(number, input)

	fmt.Printf("frequency of the digit %v is %v", input, freq)
}

func getFrequency(n int64, i int) int {
	var count int
	for n > 0 {
		temp := n % 10
		if temp == int64(i) {
			count++
		}
		n = n / 10
	}
	return count
}
