package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)    // unbuffered channel
	d := make(chan string, 5) // buffered channel

	go returnSomething("Dog", c)

	for {
		resp, open := <-c
		if !open {
			break
		}
		fmt.Println(resp)
	}

	// a better way to achive the above without an infinite for loos is
	//to use a for range loop

	go returnSomething("cat", d)

	for resp := range d {
		fmt.Println(resp)
	}

}

func returnSomething(val string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- val + " loves to play"
		time.Sleep(time.Millisecond * 500)
	}
	close(c)

}
