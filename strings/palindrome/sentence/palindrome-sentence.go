package main

import (
	"bufio"
	"fmt"
	"os"
)

func isChar(value rune) bool {
	if value >= 'A' && value <= 'Z' {
		return true
	}
	if value >= 'a' && value <= 'z' {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enater the sentence")
	scanner.Scan()
	input := scanner.Text()

	runeOfString := []rune(input)
	len := len(runeOfString)
	var isPalindrome bool = true
	i, j := 0, len-1

	for i < j {
		if !isChar(runeOfString[i]) {
			i++
			continue
		}

		if !isChar(runeOfString[j]) {
			j--
			continue
		}

		if runeOfString[i] != runeOfString[j] {
			isPalindrome = false
		}
		i++
		j--
	}

	fmt.Println(isPalindrome)

}
