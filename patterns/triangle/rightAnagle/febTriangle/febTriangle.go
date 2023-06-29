package main

import (
	"fmt"
)

func main() {
	var h int
	fmt.Println("enter the hight of the triangle")
	fmt.Scan(&h)

	var a, b int = 0, 1

	/*
	   0
	   1 1
	   2 3 5
	   8 13 21 34
	   55 89 144 233 377
	*/
	for i := 1; i <= h; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ", a)
			c := a + b
			a = b
			b = c
		}
		fmt.Println()
	}

	fmt.Println("---------")
	/*
		* * * * *
		* * * *
		* * *
		* *
		*
		above pattern in febo
	*/
	a, b = 0, 1
	fmt.Println(h)
	for i := 0; i < h; i++ {
		for j := 1; j <= h-i; j++ {
			fmt.Print(" ", a)
			c := a + b
			a = b
			b = c
		}
		fmt.Println()
	}

}
