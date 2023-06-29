package main

import (
	"fmt"
)

func main() {
	var a = 1
	var b = a

	b++

	fmt.Println("A: ", a)
	fmt.Println("B: ", b)

	var c = 1
	var d = &c

	*d++

	fmt.Println("C: ", c)
	fmt.Println("D: ", d)

}
