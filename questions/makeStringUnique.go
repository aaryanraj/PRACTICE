//Minimum deletion required to make character frequencies unique in a string

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string")
	scanner.Scan()
	input := scanner.Text()
	var str string
	unique := make(map[rune]int)
	var delRequired int

	for _, val := range []rune(input) {
		if _, found := unique[val]; found {
			unique[val]++
			delRequired++
		} else {
			unique[val] = 1
			str = str + string(val)
		}
	}
	fmt.Println(delRequired)
	fmt.Println(str)
}
