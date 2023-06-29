package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter the Number of Districts")
	scanner.Scan()
	district, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	alpha := 26 * 26
	vehPerDist := alpha * 9999
	fmt.Println("Total number of vehicals that can be registered in the state is", int(district)*vehPerDist)
}
