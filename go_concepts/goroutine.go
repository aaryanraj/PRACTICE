package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("fish")

	fmt.Scanln()
}

func count(in string) {
	for {
		fmt.Println(in)
		time.Sleep(time.Millisecond * 500)
	}
}
