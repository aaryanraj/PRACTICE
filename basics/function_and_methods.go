package main

import (
	"fmt"
)

type empDetails struct {
	name    string
	profile string
	ctc     int
}

// countdown func counts down to 0 from the given number
func countdown(num int) {
	for i := num; i > 0; i-- {
		fmt.Println(i)
	}
}

// printEmpDetails func is a method which prints employee details
func (emp *empDetails) printEmpDetails() {
	fmt.Println(emp.name)
	fmt.Println(emp.profile)
	fmt.Println(emp.ctc)
}

func main() {
	countdown(5)

	emp1 := empDetails{
		name:    "Aaryan",
		profile: "Golang Dev",
		ctc:     450000,
	}
	emp1.printEmpDetails()
}
