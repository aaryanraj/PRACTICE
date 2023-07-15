/*
Input: [[1,3],[2,6],[8,10],[15,18]]
output: [[1,6],[8,10],[15,18]]
*/

package main

import (
	"fmt"
	"sort"
)

type twoD [][]int

func (a twoD) Len() int           { return len(a) }
func (a twoD) Less(i, j int) bool { return a[i][0] < a[j][0] }
func (a twoD) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func canBeMerged(first, second []int) bool {
	if first[1] >= second[0] {
		return true
	}
	return false
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	input := [][]int{{2, 6}, {1, 3}, {5, 10}, {15, 18}}
	fmt.Println(input)
	var ans [][]int
	sort.Sort(twoD(input)) //sort on the starting time
	fmt.Println("Sorted: ", input)
	tempInterval := input[0]

	for _, interval := range input {

		if canBeMerged(tempInterval, interval) {
			tempInterval[1] = max(tempInterval[1], interval[1])

		} else {
			ans = append(ans, tempInterval)
			tempInterval = interval
		}
	}
	ans = append(ans, tempInterval)
	fmt.Println(ans)

}
