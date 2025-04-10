package main

import (
	"testing"
)

func Test_FindPairs(t *testing.T) {

	var testsSorted = []struct {
		testName            string
		inputData           []int
		target              int
		expectedPairFound   bool
		expectedFirstValue  int
		expectedSecondValue int
	}{
		{"Nil input", nil, 10, false, 0, 0},
		{"Emput input", []int{}, 10, false, 0, 0},
		{"1 input", []int{9}, 10, false, 0, 0},
		{"1 input, match", []int{9}, 9, false, 0, 0},
		{"2 input, no pair", []int{-9, 18}, 10, false, 0, 0},
		{"2 input, pair", []int{2, 8}, 10, true, 2, 8},
		{"2 input, pair", []int{2, 2, 8}, 10, true, 2, 8},
		{"3 input -ve value, pair", []int{-1, 2, 3}, 2, true, -1, 3},
		{"multiple input, pair", []int{-9, 8, 13, 19, 20, 32, 38}, 45, true, 13, 32},
		{"3 input all -ve values, pair", []int{-3, -2, -1}, -5, true, -3, -2},
	}

	for _, tt := range testsSorted {
		t.Run(tt.testName, func(t *testing.T) {
			actualPairFound, actualFirstValue, actualSecondValue := FindPairSorted(tt.inputData, tt.target)
			if actualPairFound != tt.expectedPairFound {
				t.Errorf("Pair Found:: got %v, want %v", actualPairFound, tt.expectedPairFound)
			} else if actualFirstValue != tt.expectedFirstValue {
				t.Errorf("Pair:: got %v, want %v", actualFirstValue, tt.expectedFirstValue)
			} else if actualSecondValue != tt.expectedSecondValue {
				t.Errorf("Pair:: got %v, want %v", actualSecondValue, tt.expectedSecondValue)
			}
		})

	}

	var testsUnsorted = []struct {
		testName            string
		inputData           []int
		target              int
		expectedPairFound   bool
		expectedFirstValue  int
		expectedSecondValue int
	}{
		{"Nil input", nil, 10, false, 0, 0},
		{"Emput input", []int{}, 10, false, 0, 0},
		{"1 input", []int{9}, 10, false, 0, 0},
		{"1 input, match", []int{9}, 9, false, 0, 0},
		{"2 input, no pair", []int{4, -9}, 10, false, 0, 0},
		{"2 input, pair", []int{8, 2}, 10, true, 2, 8},
		{"2 input duplicate, pair", []int{2, 8, 2}, 10, true, 8, 2},
		{"3 input -ve value, pair", []int{2, -1, 3}, 2, true, 3, -1},
		{"multiple input, pair", []int{19, -8, 1, 9, 15, 11, 4}, 16, true, 15, 1},
	}

	for _, tt := range testsUnsorted {
		t.Run(tt.testName, func(t *testing.T) {
			actualPairFound, actualFirstValue, actualSecondValue := FindPairUnSorted(tt.inputData, tt.target)
			if actualPairFound != tt.expectedPairFound {
				t.Errorf("Pair Found:: got %v, want %v", actualPairFound, tt.expectedPairFound)
			} else if actualFirstValue != tt.expectedFirstValue {
				t.Errorf("Pair:: got %v, want %v", actualFirstValue, tt.expectedFirstValue)
			} else if actualSecondValue != tt.expectedSecondValue {
				t.Errorf("Pair:: got %v, want %v", actualSecondValue, tt.expectedSecondValue)
			}
		})

	}
}
