/*
1               1
1 2           2 1
1 2 3       3 2 1
1 2 3 4   4 3 2 1
1 2 3 4 5 4 3 2 1
*/

package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println("enter the maximum number for the pattern")
	fmt.Scan(&x)
	sp := 2*x - 3
	for i := 1; i <= x; i++ {
		for j := 1; j <= x; j++ {
			if j <= i {
				fmt.Print(j, " ")
			}
		}
		for j := 1; j <= sp; j++ {
			fmt.Print("  ")
		}
		for j := i; j >= 1; j-- {
			if j != x {
				fmt.Print(j, " ")
			}
		}
		fmt.Println()
		sp -= 2
	}
}
