// check if the string is palindrome using recusion

package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(sl []rune, i int) bool {
	n := len(sl)
	if i >= n/2 {
		return true
	}
	if sl[i] != sl[n-i-1] {
		return false
	}
	return isPalindrome(sl, i+1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string")
	scanner.Scan()
	input := scanner.Text()

	slOfInput := []rune(input)
	ans := isPalindrome(slOfInput, 0)
	if ans {
		fmt.Println("The string is Plaindrome")
	} else {
		fmt.Println("The string is not Plindrome")
	}
}
