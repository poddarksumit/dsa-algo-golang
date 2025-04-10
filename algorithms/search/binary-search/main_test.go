package main

import "testing"

func Test_BinarySearch(tt *testing.T) {

	var tests = []struct {
		testname              string
		numbers               []int
		target                int
		expectedIndexOfTarget int
	}{
		{"Empty numbers", nil, 5, -1},
		{"1 element, not found", []int{3}, 5, -1},
		{"1 element, found", []int{3}, 3, 0},
		{"Index not found", []int{3, 7, 10, 15, 25, 56}, 5, -1},
		{"Index found", []int{3, 7, 10, 15, 25, 56}, 15, 3},
	}

	for _, test := range tests {
		tt.Run(test.testname, func(t *testing.T) {
			ans := binarySearch(test.numbers, test.target)

			if ans != test.expectedIndexOfTarget {
				tt.Error("Expected %i, got %i \n", test.expectedIndexOfTarget, ans)
			}
		})
	}
}
