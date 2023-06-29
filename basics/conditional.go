package main

import (
	"fmt"
)

func main() {
	age := 27

	if age >= 21 {
		fmt.Println("You are allowed here")
	} else {
		fmt.Println("Come back when you turn 21")
	}

	age = 19
	gender := "Female"
	if age < 21 && gender == "Male" {
		fmt.Println("Well You are Not old enough to get Marrid")
	} else if age >= 18 && gender == "Female" {
		fmt.Println("You are old enough to get Marrid")
	} else if age >= 21 && gender == "Male" {
		fmt.Println("You are old enough to get Marrid")
	} else if age < 18 {
		fmt.Println("You are old enough to get Marrid")
	} else {
		fmt.Println("Case not valid")
	}
}
