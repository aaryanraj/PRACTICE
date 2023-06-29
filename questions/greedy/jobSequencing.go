/*
You are given a set of N jobs where each job comes with a deadline and profit. The profit can only be earned
upon completing the job within its deadline. Find the number of jobs done and the maximum profit that can be
obtained. Each job takes a single unit of time and only one job can be performed at a time.
Input: N = 4, Jobs = {(1,4,20),(2,1,10),(3,1,40),(4,1,30)}
Output: 2 60
*/

package main

import (
	"fmt"
	"sort"
)

type job struct {
	index    int
	deadline int
	value    int
}

func pickJobs(jobs []job, n, slots int) (int, int) {
	sArr := make([]int, slots+1)
	for i := 1; i <= slots; i++ {
		sArr[i] = -1
	}
	var jobCount, jobProfit int
	for i := 0; i < n; i++ {
		for j := jobs[i].deadline; j > 0; j-- {
			if sArr[j] == -1 {
				sArr[j] = i
				jobCount++
				jobProfit = jobProfit + jobs[i].value
				break
			}
		}
	}
	return jobCount, jobProfit
}

func main() {
	jobs := [][]int{{1, 2, 100}, {2, 1, 19}, {3, 2, 27}, {4, 1, 25}, {5, 1, 15}}
	slots, n := 0, 5
	var jobValue []job

	for _, value := range jobs {
		jobValue = append(jobValue, job{index: value[0], deadline: value[1], value: value[2]})
		if value[2] > slots {
			slots = value[2]
		}
	}
	sort.Slice(jobValue, func(i, j int) bool {
		return jobValue[i].value > jobValue[j].value
	})
	num, profit := pickJobs(jobValue, n, slots)

	fmt.Println(num, profit)

}
