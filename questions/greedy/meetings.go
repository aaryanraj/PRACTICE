//N meetings in one room
/*There is one meeting room in a firm. You are given two arrays, start and end each of size N.
For an index ‘i’, start[i] denotes the starting time of the ith meeting while end[i]  will denote the ending
time of the ith meeting. Find the maximum number of meetings that can be accommodated if only one meeting can
happen in the room at a  particular time. Print the order in which these meetings will be performed.
Input:  N = 6,  start[] = {1,3,0,5,8,5}, end[] =  {2,4,5,7,9,9}
Output: 1 2 4 5
*/
package main

import (
	"fmt"
	"sort"
)

type meeting struct {
	start int
	end   int
	pos   int
}

func main() {
	var n int = 6
	var start, end = []int{1, 3, 0, 5, 8, 5}, []int{2, 4, 5, 7, 9, 9}
	var ans []int
	var meet []meeting

	for i := 0; i < n; i++ {
		m := meeting{start: start[i], end: end[i], pos: i + 1}
		meet = append(meet, m)
	}
	sort.SliceStable(meet, func(i, j int) bool {
		return meet[i].end < meet[j].end
	})
	ans = append(ans, meet[0].pos)
	limit := meet[0].end

	for i := 1; i < n; i++ {
		if meet[i].start > limit {
			ans = append(ans, meet[i].pos)
			limit = meet[i].end
		}
	}
	fmt.Println(ans)
}
