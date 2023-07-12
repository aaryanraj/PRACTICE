package main

import (
	"fmt"
	"time"
)

func countAnimal(in string) {
	for {
		fmt.Println(in)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go countAnimal("sheep")
	go countAnimal("fish")

	fmt.Scanln()
}
