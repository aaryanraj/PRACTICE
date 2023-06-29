package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "every 200ms"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)

	}
}

// in the above example we are waiting for the channel1 to transmit data every 500ms
// this make sthe faster go-routine also wait for 500ms since the channels are blocking in nature
// to overcome this problem we can use select statement,
// select will not wait for a channel syncronusly, it will get the data from channels randomly if it has a data.
