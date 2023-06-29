package main

import (
	"fmt"
)

func main() {
	var n1, n2 int64
	var base int
	fmt.Println("Enter the first number")
	fmt.Scan(&n1)

	fmt.Println("Enter the second number")
	fmt.Scan(&n2)

	fmt.Println("Enter the base of the above numbers")
	fmt.Scan(&base)

	ans := anyBaseAddition(n1, n2, base)
	fmt.Println(ans)
}

func anyBaseAddition(n1 int64, n2 int64, base int) int64 {
	var mul int64 = 1
	var carry, sum int = 0, 0
	var final int64 = 0
	for n1 > 0 || n2 > 0 || carry > 0 {
		digit1 := int(n1 % 10)
		digit2 := int(n2 % 10)

		sum = digit1 + digit2 + carry
		carry = sum / base
		value := sum % base
		final += mul * int64(value)
		mul = mul * 10

		n1 = n1 / 10
		n2 = n2 / 10
	}

	return final
}
