package main

import (
	"fmt"
)

func main() {
	var num int64
	var oldBase, newBase int
	fmt.Println("Enter the number")
	fmt.Scan(&num)
	fmt.Println("Enter the base of the above number")
	fmt.Scan(&oldBase)
	fmt.Println("Enter the new base")
	fmt.Scan(&newBase)

	dec := anyToDecimal(num, oldBase)
	ans := decimalToAny(dec, newBase)
	fmt.Println(ans)
}

func anyToDecimal(num int64, oldBase int) int64 {
	var ans, mul int64 = 0, 1
	for num > 0 {
		rem := num % 10
		num = num / 10

		ans = ans + rem*mul
		mul *= int64(oldBase)
	}
	return ans
}

func decimalToAny(num int64, newBase int) int64 {
	var ans, mul int64 = 0, 1

	for num > 0 {
		rem := num % int64(newBase)
		num = num / int64(newBase)

		ans = ans + rem*mul
		mul = mul * 10
	}
	return ans
}
