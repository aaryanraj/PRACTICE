package main

import (
	"fmt"
)

func main() {
	var dec int64
	var base int
	fmt.Println("Enter the decimal number")
	fmt.Scan(&dec)
	fmt.Println("Enter the base")
	fmt.Scan(&base)

	res := anyBaseConverter(dec, base)
	fmt.Printf("Decimal number %v is equal to %v in the base of %v", dec, res, base)
}

// to get the converted number long devide the decimal number with the base
// the number in the other base will be r1*10^0 + r2*10^1 + r3*10^2 .....
func anyBaseConverter(dec int64, base int) int64 {
	var ans, mul int64 = 0, 1
	for dec > 0 {
		rem := dec % int64(base)
		dec = dec / int64(base)

		ans = ans + rem*mul
		mul *= 10
	}
	return ans
}
