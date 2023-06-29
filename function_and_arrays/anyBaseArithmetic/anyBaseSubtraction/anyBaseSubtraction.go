package main

import "fmt"

func main() {
	var n1, n2 int64
	var base int
	fmt.Println("Enter the first number")
	fmt.Scan(&n1)

	fmt.Println("Enter the second number")
	fmt.Scan(&n2)

	fmt.Println("Enter the base of the above numbers")
	fmt.Scan(&base)
	if n1 > n2 {
		ans := anyBaseSubtraction(n1, n2, base)
		fmt.Println(ans)
	} else {
		ans := anyBaseSubtraction(n2, n1, base)
		fmt.Printf("-%v", ans)
	}
}

func anyBaseSubtraction(n1 int64, n2 int64, base int) int64 {
	var final, mul int64 = 0, 1
	var borrow int = 0

	for n1 > 0 {
		digit1 := n1 % 10
		digit2 := n2 % 10

		n1 = n1 / 10
		n2 = n2 / 10

		dif := digit1 - digit2 - int64(borrow)

		if dif < 0 {
			borrow = 1
			dif = dif + int64(base)
		} else {
			borrow = 0
		}

		final = final + (mul * dif)
		mul = mul * 10
	}
	return final
}
