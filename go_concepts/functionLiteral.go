//This example is just to build background for closure function, check that as well
/*Go language provides a special feature known as an anonymous function. An anonymous function can form a
closure. A closure is a special type of anonymous function that references variables declared outside of
the function itself. It is similar to accessing global variables which are available before the declaration
of the function.*/

package main

import (
	"fmt"
)

func main() {
	x := 0

	counter := func() int {
		x++
		return x
	}

	fmt.Println(counter)
	fmt.Println(counter)

}

/*
Explanation: The variable GFG was not passed as a parameter to the anonymous function but the function has
access to it. In this example, there is a slight problem as any other function which will be defined in the
main has access to the global variable GFG and it can change it without calling counter function.
Thus closure also provides another aspect which is data isolation
*/
