package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork() int {
	time.Sleep(time.Second)
	return rand.Int()
}
func main() {

	dataChan := make(chan int)
	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				resp := doWork()
				dataChan <- resp
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for val := range dataChan {
		fmt.Println(val)
	}
}
