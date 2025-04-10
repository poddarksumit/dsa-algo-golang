package main

import "fmt"

func main() {

	fmt.Println("Welcome to Binary Search Program!!")

	numbers := []int{1, 9, 17, 24, 45, 67, 82, 190}

	indexOfTarget := binarySearch(numbers, 9)
	fmt.Println("Index of target: ", indexOfTarget)
}

func binarySearch(numbers []int, target int) int {

	leftIndex, rightIndex := 0, len(numbers)-1

	for leftIndex <= rightIndex {
		midIndex := (leftIndex + rightIndex) / 2

		if numbers[midIndex] == target {
			return midIndex
		} else if numbers[midIndex] < target {
			leftIndex = midIndex + 1
		} else if numbers[midIndex] > target {
			rightIndex = midIndex - 1
		}
	}

	return -1
}
