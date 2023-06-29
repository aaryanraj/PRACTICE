package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string with no spaces")
	scanner.Scan()
	input := scanner.Text()

	runeOfString := []rune(input)
	length := len(runeOfString)
	var isPalindrome bool = true
	i := 0
	j := length - 1
	for i < j {
		if runeOfString[i] != runeOfString[j] {
			isPalindrome = false
		}
		i++
		j--
	}
	fmt.Println(isPalindrome)
}
