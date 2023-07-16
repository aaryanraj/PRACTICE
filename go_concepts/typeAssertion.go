package main

import (
	"fmt"
)

func main() {
	var myVar interface{} = 7677

	if str, ok := myVar.(string); ok {
		fmt.Println(str)
	} else {
		fmt.Println("myVar is not a string")
	}

	switch myVar.(type) {
	case int:
		fmt.Println("myVar is an integer")
	case string:
		fmt.Println(",yVar is an string")
	case float64:
		fmt.Println("myVar is a float64")
	case bool:
		fmt.Println("myVar is a boolean")
	default:
		fmt.Println("I dont know the type of myVar")
	}
}
