package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		countSheep("sheep")
		//count("dog")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		count("dog")
		wg.Done()
	}()
	wg.Wait()
}

func countSheep(thing string) {
	for i := 0; i < 5; i++ {
		fmt.Println(thing)
	}
}

func count(thing string) {
	//time.Sleep(time.Microsecond)
	for i := 0; i < 10; i++ {
		fmt.Println(thing)
	}
}
