package main

import (
	"fmt"
)

func main() {
	var (
		num1       = 2
		num2       = 3
		sum        int
		difference int
		product    int
		remainder  float32
		quotient   int
	)
	// Arithemetic Operators
	sum = num1 + num2
	difference = num1 - num2
	product = num1 * num2
	remainder = float32(num1 % num2)
	quotient = num1 / num2

	fmt.Println("Sum: ", sum)
	fmt.Println("Difference: ", difference)
	fmt.Println("Product: ", product)
	fmt.Println("Remainder: ", remainder)
	fmt.Println("quotient: ", quotient)

	// Relational Operators
	var result bool

	result = num1 > num2 // greater than
	fmt.Println("is Num1 Greater than Num2: ", result)
	result = num1 < num2 // less than
	fmt.Println("is Num1 Less than Num2: ", result)
	result = num1 >= num2 // greater than or equal to
	fmt.Println("is Num1 Greater or equal to Num2: ", result)
	result = num1 <= num2 // less than or equal to
	fmt.Println("is Num1 less or equal to than Num2: ", result)
	result = num1 == num2 // equals to
	fmt.Println("is Num1 equals to Num2: ", result)
	result = num1 != num2 // not equals to
	fmt.Println("is Num1 Not equals to Num2: ", result)

	// Logical operator
	const lastName = "Raj"
	var age = 20

	var inviteToParty = lastName == "Sharma" || lastName == "Raj" && age > 18
	fmt.Println(fmt.Sprintf("Abhishek %s is invited to party: ", lastName), inviteToParty)
	var isAllowedToDrink = lastName == "Sharma" && (age >= 21 && age < 60)
	fmt.Println("Abhishek Sharma allowed to drink: ", isAllowedToDrink)

}
