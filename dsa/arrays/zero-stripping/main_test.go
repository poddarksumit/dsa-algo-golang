package main

import (
	"testing"
)

func Test_ZeroStripping(t *testing.T) {

	var tests = []struct {
		testName       string
		inputData      [][]int
		expectedOutput [][]int
	}{
		{"Nil input", nil, nil},
		{"Empty input", [][]int{}, [][]int{}},
		{"With input", [][]int{
			{0, 11, 14},
			{14, 31, 10},
			{90, 71, 0},
		}, [][]int{
			{0, 0, 0},
			{0, 31, 0},
			{0, 0, 0},
		}},
	}

	for _, tt := range tests {

		t.Run(tt.testName, func(test *testing.T) {

			actualOutput := StripZero(tt.inputData)

			for row := range actualOutput {
				for col := range len(actualOutput[row]) {

					if actualOutput[row][col] != tt.expectedOutput[row][col] {
						t.Errorf("Got mismatch for %v: got %v, expected %v", tt.testName, actualOutput[row][col], tt.expectedOutput[row][col])
					}
				}
			}
		})
	}
}
