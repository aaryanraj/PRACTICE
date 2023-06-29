package main

import "fmt"

func main() {
	var favBook = "Rashmirathi"
	//other ways
	//favBook:= "Rashmirathi"

	//print in new line
	fmt.Println(favBook)

	//reassign
	favBook = "Krishnavatar"
	fmt.Println(favBook)

	//explicit defining of type
	var favMovie string = "there are many"
	fmt.Println(favMovie)

	var name string
	var age int
	var amiavalible bool

	//printing with the default values
	fmt.Println(fmt.Sprintf("Hello my name is %s I am %d years old and i am currently available %v", name, age, amiavalible))

	//assigning values to the variables
	name = "Aaryan"
	age = 27
	amiavalible = true

	//printing with the assigned values
	fmt.Println(fmt.Sprintf("Hello my name is %s I am %d years old and i am currently available %v", name, age, amiavalible))

	//Block creation

	var (
		favCricketer = "Sachin"
		valueOfPi    = 3.14
		favIPLTeam   = "CSK"
	)

	fmt.Println(favCricketer, valueOfPi, favIPLTeam)

	//constants
	const MyName = "Aaryan"
	// we cannot change the value of a constant
	// MyName = "Shubham"  is not a valid statement anymore

	//Declaring multiple variable with shortcut
	pet1, pet2 := "Dog", "Cat"

	fmt.Println("Pet 1 :", pet1, "  ", "Pet 2 :", pet2)
}
