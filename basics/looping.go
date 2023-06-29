package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// the above code breaks down to this
	k := 1
	for {
		fmt.Print(k, " ")
		if k == 5 {
			break
		}
		k++
	}
	fmt.Println()
	//continue skips the iteration
	r := 1
	for r < 50 {
		if r%2 == 0 {
			r++
			continue
		}
		fmt.Print(r, " ")
		r++
	}
}
