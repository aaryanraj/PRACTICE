//Print from 1 to n in backTracking fashion.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printNumbers(n int) {
	if n < 1 {
		return
	}
	printNumbers(n - 1)
	fmt.Println(n)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number n")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	printNumbers(int(input))
}
