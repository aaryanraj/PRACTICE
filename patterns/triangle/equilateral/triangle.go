package main

import (
	"fmt"
)

func main() {
	var l int
	fmt.Println("Enter the length of the triangle")
	fmt.Scan(&l)

	for i := 1; i <= l; i++ {
		for j := 1; j <= l-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
