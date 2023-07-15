package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("hello")
	}()

	go func() {
		fmt.Println("world")
	}()

	fmt.Println("Finish")

}
