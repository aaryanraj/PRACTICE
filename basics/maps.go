package main

import "fmt"

func main() {
	nameAndCost := make(map[string]int)

	nameAndCost["soap"] = 30
	nameAndCost["shampoo"] = 60
	nameAndCost["cream"] = 100
	/*
		the above can also be written as

		nameAndCost:= map[string]int{
			"soap":30,
			"shampoo":60,
			"cream":100,
		}
	*/

	nameAndCost["facewash"] = 120
	delete(nameAndCost, "soap")

	fmt.Println(nameAndCost)

	for key, value := range nameAndCost {

		if key == "shanpoo" {
			fmt.Println(key, value)
		}
		if value == 120 {
			fmt.Println(nameAndCost[key])
		}
	}

}
