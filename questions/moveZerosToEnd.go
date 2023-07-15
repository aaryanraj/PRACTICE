// move all the zeros in the array to right side
// first approch -- swap when we find a non zero element in a for loop
// second approch -- check for zeros and delete them and add the zeros at last
// third approch -- save all the non zero in one array and add zeros at last
package main

import "fmt"

//first approch
/*
func swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
func main() {
	input := []int{5, 0, 2, 0, 7, 0, 4}
	j := 0
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			swap(input, j, i)
			j++
		}
	}
	fmt.Println(input)
}
*/
/*
func main() {
	input := []int{0, 1, 0, 0, 0, 0, 4}
	len := len(input)
	for i := 0; i < len; i++ {
		if input[i] == 0 {
			input = append(input[:i], input[i+1:]...)
			input = append(input, 0)
			len--
			i--
		}
	}
	fmt.Println(input)
}
*/
func main() {
	input := []int{0, 1, 0, 0, 0, 0, 4}
	var zero, nonZero []int
	len := len(input)
	for i := 0; i < len; i++ {
		if input[i] == 0 {
			zero = append(zero, input[i])
		} else {
			nonZero = append(nonZero, input[i])
		}
	}
	nonZero = append(nonZero, zero...)
	fmt.Println(nonZero)
}
