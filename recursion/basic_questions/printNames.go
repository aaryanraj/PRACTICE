// Print your name 5 time using recursion

package main

import (
	"fmt"
)

var count = 0

func printName() {
	if count == 5 {
		return
	}
	fmt.Println("Aaryan")
	count++
	printName()
}
func main() {
	printName()
}
