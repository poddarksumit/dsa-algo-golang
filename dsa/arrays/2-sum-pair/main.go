package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to finding pair to sum")

	input := []int{-1, 2, 4, 8, 10}
	target := 10

	foundPair, firstValue, secondValue := FindPairSorted(input, target)

	if foundPair {
		fmt.Println("Pair found: ", firstValue, secondValue)
	} else {
		fmt.Println("No Pair found!")
	}

	input = []int{-18, 12, 4, 3, 10, 19, 8, -3}
	target = 7

	foundPair, firstValue, secondValue = FindPairUnSorted(input, target)

	if foundPair {
		fmt.Println("Unsorted Pair found: ", firstValue, secondValue)
	} else {
		fmt.Println("Unsorted No Pair found!")
	}
}

func FindPairSorted(input []int, target int) (bool, int, int) {

	firstIndex, lastIndex := 0, len(input)-1

	for firstIndex < lastIndex {

		sum := input[firstIndex] + input[lastIndex]

		if sum < target {
			firstIndex++
		} else if sum > target {
			lastIndex--
		} else {
			return true, input[firstIndex], input[lastIndex]
		}
	}

	return false, 0, 0
}

func FindPairUnSorted(input []int, target int) (bool, int, int) {

	numberIndexMap := map[int]int{}

	for inputIndex, inputValue := range input {

		numberIndexMap[inputValue] = inputIndex
		numberToSearch := target - inputValue

		value, exists := numberIndexMap[numberToSearch]

		if exists {
			return true, input[inputIndex], input[value]
		}
	}

	return false, 0, 0
}
