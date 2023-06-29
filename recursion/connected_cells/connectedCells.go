package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func findMaxConnections(matrix [][]int, m, n int) int {
	maxConections := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			maxConections = int(math.Max(float64(maxConections), findConnections(matrix, m, n, r, c)))
		}
	}
	return maxConections
}

func findConnections(matrix [][]int, m, n, r, c int) float64 {
	var ans float64 = 0

	if r < 0 || c < 0 || r >= m || c >= n {
		ans = 0
	} else if matrix[r][c] == 1 {
		matrix[r][c] = 0
		ans = 1 +
			findConnections(matrix, m, n, r-1, c) +
			findConnections(matrix, m, n, r+1, c) +
			findConnections(matrix, m, n, r, c-1) +
			findConnections(matrix, m, n, r, c+1) +
			findConnections(matrix, m, n, r-1, c-1) +
			findConnections(matrix, m, n, r-1, c+1) +
			findConnections(matrix, m, n, r+1, c-1) +
			findConnections(matrix, m, n, r+1, c+1)
	}
	return ans
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	matrix := make([][]int, m)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sc.Scan()
			v, _ := strconv.Atoi(sc.Text())
			matrix[i][j] = v
		}
	}
	fmt.Println(matrix)

	fmt.Println(findMaxConnections(matrix, m, n))
}
