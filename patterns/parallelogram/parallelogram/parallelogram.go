package main

import (
	"fmt"
)

func main() {
	var l, b int
	fmt.Println("enter the length of the ||gm")
	fmt.Scan(&l)
	fmt.Println("enter the breadth of the ||gm")
	fmt.Scan(&b)

	for i := 1; i <= b; i++ {
		for j := 1; j <= b-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= l; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
