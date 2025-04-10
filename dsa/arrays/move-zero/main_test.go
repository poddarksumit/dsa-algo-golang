package main

import "testing"

func Test_MoveZero(t *testing.T) {

	var tests = []struct {
		testName     string
		inputData    []int
		expectedData []int
	}{
		{"Nil input", nil, nil},
		{"Emput input", []int{}, []int{}},
		{"1 input", []int{9}, []int{9}},
		{"2 input, no zero", []int{9, 7}, []int{9, 7}},
		{"2 input, zero", []int{0, 9}, []int{9, 0}},
		{"3 input, zero", []int{0, 0, 0}, []int{0, 0, 0}},
		{"3 input, no zero", []int{1, 3, 2}, []int{1, 3, 2}},
		{"input, with right zero", []int{1, 1, 1, 0, 0}, []int{1, 1, 1, 0, 0}},
		{"input input, with left zero", []int{0, 0, 1, 1, 1}, []int{1, 1, 1, 0, 0}},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			actualInput := moveZero(test.inputData)

			for index, value := range actualInput {
				if value != test.expectedData[index] {
					t.Errorf("Data not matched: got %v, want %v", actualInput, test.expectedData)
				}
			}
		})
	}

}
