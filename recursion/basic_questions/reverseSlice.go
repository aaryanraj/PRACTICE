// reverse the given slice usingt recursion.

package main

import (
	"fmt"
)

func reverseSl(sl []int, l, r int) []int {
	if l >= r {
		return sl
	}
	temp := sl[l]
	sl[l] = sl[r]
	sl[r] = temp
	return reverseSl(sl, l+1, r-1)
}

func main() {
	sl := []int{9, 5, 7, 3, 2, 6}
	l, r := 0, len(sl)-1
	revSl := reverseSl(sl, l, r)
	fmt.Println(revSl)
}
