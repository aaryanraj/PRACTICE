package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter length of the triangle")
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
