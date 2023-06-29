package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number")
	sc.Scan()
	input, _ := strconv.ParseInt(sc.Text(), 10, 32)
	var output string
	if input%3 == 0 {
		output = output + "Fizz"
	}
	if input%7 == 0 {
		output = output + "Buzz"
	}
	if output == "" {
		output = strconv.Itoa(int(input))
	}

	fmt.Println(output)

}
