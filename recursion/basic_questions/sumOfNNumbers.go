// Print the sum of first n numbers using recursion.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumOfNNumbers(n, sum int) {
	if n < 1 {
		fmt.Println(sum)
		return
	}
	sum = sum + n
	sumOfNNumbers(n-1, sum)
}

func total(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + total(n-1)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number n")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	sum := 0
	sumOfNNumbers(int(input), sum)
	answer := total(int(input))
	fmt.Println(answer)
}
