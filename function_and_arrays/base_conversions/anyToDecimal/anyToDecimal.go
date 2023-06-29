package main

import (
	"fmt"
)

func main() {
	var any int64
	var base int
	fmt.Println("Enter the number in any base")
	fmt.Scan(&any)
	fmt.Println("Enter the base of the number")
	fmt.Scan(&base)
	ans := anyToDecimal(any, base)
	fmt.Println(ans)
}

// to convert any base to decimal, long divide the given number by 10
// the decimal number will be r1*base^0 + r2*base^1 + r3*base^2 ....
func anyToDecimal(any int64, base int) int64 {
	var ans, mul int64 = 0, 1
	for any > 0 {
		rem := any % 10
		any = any / 10

		ans = ans + rem*mul
		mul *= int64(base)
	}
	return ans
}
