package main

import "fmt"

func main() {
	fmt.Println("Welcome to Move Zero!")

	newInput := moveZero([]int{1, 0, 3, 2, 5, 0})

	fmt.Println(newInput)
}

func moveZero(input []int) []int {

	if len(input) == 0 {
		return input
	}

	zeroIndex := 0

	for nonZeroIndex, value := range input {
		if value != 0 {
			zeroIndexValue := input[zeroIndex]
			input[zeroIndex] = value
			input[nonZeroIndex] = zeroIndexValue
			zeroIndex++
		}
	}

	return input
}
