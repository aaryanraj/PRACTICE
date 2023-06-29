//Given a vehical registration number of the format JH01EW7845 find the rank of the vehical in the district.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the vehical registration number")
	scanner.Scan()
	input := scanner.Text()
	runeOfInput := []rune(input)
	alpha := runeOfInput[4:6]
	tail := runeOfInput[6:]
	intTail, _ := strconv.ParseInt(string(tail), 10, 32)
	var countAlpha int
out:
	for i := 'a'; i <= 'z'; i++ {
		for j := 'a'; j <= 'z'; j++ {
			countAlpha++
			if i == alpha[0] && j == alpha[1] {
				break out
			}
		}
	}
	fmt.Println("The Rank of your vehical in the district is", ((countAlpha-1)*9999)+int(intTail))

}
