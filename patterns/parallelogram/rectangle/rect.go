package main

import (
	"fmt"
)

func main() {
	var l, b int
	fmt.Println("enter the length of the rect")
	fmt.Scan(&l)
	fmt.Println("enter the breadth of the rect")
	fmt.Scan(&b)

	for i := 0; i < b; i++ {

		for j := 0; j < l; j++ {
			fmt.Print("* ")
		}
		fmt.Println()

	}
}
