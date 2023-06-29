//print from n to 1 in backtracking fashion

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func numberPrinter(count, n int) {
	if count > n {
		return
	}
	numberPrinter(count+1, n)
	fmt.Println(count)

}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number n")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	numberPrinter(1, int(input))
}
