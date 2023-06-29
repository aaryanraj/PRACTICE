package main

import (
	"fmt"
)

func main() {
	city := "Delhi"

	switch city {
	case "Bengaluru":
		fmt.Println("You Must be Somehow related to IT")
	case "Delhi":
		fmt.Println("How is the Air quality doing")
	case "Kolkata":
		fmt.Println("Is it really worth living")
	case "Chennai":
		fmt.Println("How's the heat and humidity")
	case "Mumbai":
		fmt.Println("Do you have your own house!!")
	default:
		fmt.Println("Life is good!! Metros are over srated")
	}
}
