// Print numbers from 1 to N using recursion
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printFrom1ToN(n, count int) {
	if count > n {
		return
	}
	fmt.Println(count)
	count++
	printFrom1ToN(n, count)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number N")
	scanner.Scan()
	var count int = 1
	input, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	printFrom1ToN(int(input), count)
}
