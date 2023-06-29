package main

import (
	"fmt"
)

func towerOfHanoi(n int, from, to, temp string) {
	if n == 1 {
		fmt.Println("Move disk", n, "from tower", from, "to tower", to)
		return
	}

	towerOfHanoi(n-1, from, temp, to)

	fmt.Println("Move disk", n, "from tower", from, "to tower", to)

	towerOfHanoi(n-1, temp, to, from)
}

func toh(n int) {
	towerOfHanoi(n, "A", "C", "B")
}

func main() {
	toh(4)
}
