package main

import "fmt"

func main() {
	fmt.Println("Welcome to Zero Stripping")

	input := [][]int{
		{0, 11, 14},
		{14, 31, 10},
		{90, 71, 0},
	}

	input = StripZero(input)

	fmt.Println(input)
}

func StripZero(input [][]int) [][]int {

	if input == nil {
		return input
	}

	rowSet, columnSet := make(map[int]bool), make(map[int]bool)

	for rowIndex := range input {
		for columnIndex := range len(input[rowIndex]) {

			if input[rowIndex][columnIndex] == 0 {
				rowSet[rowIndex] = true
				columnSet[columnIndex] = true
			}
		}
	}

	for rowIndex := range input {
		for columnIndex := range len(input[rowIndex]) {

			rowIndexExists := rowSet[rowIndex]
			columnIndexExists := columnSet[columnIndex]

			if rowIndexExists || columnIndexExists {
				input[rowIndex][columnIndex] = 0
			}
		}
	}

	return input
}
