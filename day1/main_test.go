package main

import "testing"

type calibrateTest struct {
	input          []int
	expectedResult int
}

func TestCalibrate(t *testing.T) {
	var tests = []calibrateTest{
		{[]int{1, -2, 3, 1}, 3},
		{[]int{1, 1, 1}, 3},
		{[]int{-1, -2, -3}, -6},
	}

	for _, test := range tests {
		result := calibrate(0, test.input)
		if result != test.expectedResult {
			t.Fatalf("Calibrate returned wrong result, expecting %d got %d", test.expectedResult, result)
		}
	}

}
