//Minimum number of platforms required for a railway
/*We are given two arrays that represent the arrival and departure times of trains that stop at the platform.
 We need to find the minimum number of platforms needed at the railway station so that no train has to wait
 Input: N=6,
arr[] = {9:00, 9:45, 9:55, 11:00, 15:00, 18:00}
dep[] = {9:20, 12:00, 11:30, 11:50, 19:00, 20:00}
Output:3
*/

package main

import (
	"fmt"
	"sort"
)

func platformRequired(arr, dep []int, n int) int {
	sort.Ints(arr)
	sort.Ints(dep)
	platOccupied, platReq := 1, 1

	for i, j := 1, 0; j < n && i < n; {
		if arr[i] <= dep[j] {
			platOccupied++
			i++

		} else if arr[i] > dep[j] {
			platOccupied--
			j++
		}

		if platOccupied > platReq {
			platReq = platOccupied
		}
	}
	return platReq
}

func main() {
	arr := []int{900, 945, 955, 1100, 1500, 1800}
	dep := []int{920, 1200, 1130, 1150, 1900, 2000}
	len := len(arr)

	platReq := platformRequired(arr, dep, len)
	fmt.Println(platReq)
}
