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

func TestCalibratePart2(t *testing.T) {
	var tests = []calibrateTest{
		{[]int{1, -2, 3, 1}, 2},
		{[]int{1, -1}, 0},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, 3, 8, 5, -6}, 5},
		{[]int{7, 7, -2, -7, -4}, 14},
	}

	for _, test := range tests {
		result := calibratePart2(0, test.input)
		if result != test.expectedResult {
			t.Fatalf("CalibratePart2 returned wrong result, expecting %d got %d", test.expectedResult, result)
		}
	}

}
