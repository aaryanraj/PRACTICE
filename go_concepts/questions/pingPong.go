package main

import (
	"fmt"
	"time"
)

func playerOne(ch1, ch2 chan int) {
	for {
		hit := <-ch1
		fmt.Println("Player 1 hit ", hit)
		time.Sleep(500 * time.Millisecond)
		ch2 <- hit
	}
}

func playerTwo(ch1, ch2 chan int) {
	for {
		hit := <-ch2
		fmt.Println("Player 2 hit ", hit)
		time.Sleep(500 * time.Millisecond)
		hit++
		ch1 <- hit
	}
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go playerOne(ch1, ch2)
	go playerTwo(ch1, ch2)

	ch1 <- 1
	select {}

}
