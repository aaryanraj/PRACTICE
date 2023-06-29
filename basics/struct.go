package main

import (
	"fmt"
)

type Animal struct {
	species  string
	lifeSpan int
	petable  bool
}

func main() {
	var animal1 = Animal{
		species:  "Dog",
		lifeSpan: 12,
		petable:  true,
	}

	fmt.Println(animal1)

	animal1.lifeSpan = 13
	animal1.species = "cat"

	fmt.Println(animal1)

	animal2 := Animal{"Cow", 15, true}
	fmt.Println(animal2)

	animal3 := Animal{}
	fmt.Println(animal3)
	animal3.petable = false
	animal3.lifeSpan = 23
	animal3.species = "Lion"
	fmt.Println(animal3)
}
