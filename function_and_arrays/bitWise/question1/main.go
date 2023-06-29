// Take an integer input from the user and print XOR of all the numbers between 1 to n

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter the value of N")
	fmt.Scan(&n)
	if n%4 == 0 {
		fmt.Println(n)
	} else if n%4 == 1 {
		fmt.Println("1")
	} else if n%4 == 2 {
		fmt.Println(n + 1)
	} else {
		fmt.Println(0)
	}
}
