package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Arrays
	purchases := [5]float32{19.99, 20.67, 1.50, 15.99}
	income := [4]float32{2034.56, 4245.90, 5098.43}
	sales := [...]float32{}

	income[3] = 6789.00
	fmt.Println(purchases, "  ", reflect.TypeOf(purchases))
	fmt.Println(income)
	fmt.Println(sales)

	for i := 0; i < len(purchases); i++ {
		fmt.Println(purchases[i])
	}

	//Slices

	mySlice := purchases[2:] // abstraction over the array purcheses
	fmt.Println(mySlice, "  ", reflect.TypeOf(mySlice))

	//another way
	mySl := []string{}
	mySl1 := []string{"test", "something", "normal"}

	fmt.Println(mySl)
	mySl = append(mySl1, "just", "for", "fun")
	fmt.Println(mySl)

	//another way
	sl := make([]int, 5, 5) // parameter are type,lenght,capacity
	// sl:= make([]int,5) is also valid, here we skip capacity
	fmt.Println(sl) // we will get a slice of 5 elements initialised to 0

	//appending two Slices
	s1 := []string{"Aaryan"}
	appendedSl := append(s1, mySl1...)
	fmt.Println(appendedSl)
}
