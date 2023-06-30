// print the nth fibonacchi number of seriese using recusrion

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fibonacchi(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacchi(n-1) + fibonacchi(n-2)
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter the number")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	//0,1,1,2,3,5,8.....
	//considering 0 as the 0th and 1 as the first fibonacchi numbers
	ans := fibonacchi(int(input))
	fmt.Println(ans)
}
