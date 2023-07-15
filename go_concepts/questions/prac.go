package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		helloworld()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		goodbye()
	}()

	wg.Wait()

}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {

	fmt.Println("Good Bye!")
}
