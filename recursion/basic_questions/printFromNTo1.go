package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printFromNTO1(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	n--
	printFromNTO1(n)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the nymber N")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 32)

	printFromNTO1(int(input))

}
