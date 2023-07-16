// check out functionLiteral example before trying to understand this.
package main

import (
	"fmt"
)

func newCounter() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	counter := newCounter()

	fmt.Println(counter())
	fmt.Println(counter())
}

/*
Explanation: The closure references the variable GFG even after the newCounter() function has finished
running but no other code outside of the newCounter() function has access to this variable. This is how data
 persistency is maintained between function calls while also isolating the data from other code.
*/
