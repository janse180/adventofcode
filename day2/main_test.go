package main

import (
	"reflect"
	"testing"
)

type checksumTest struct {
	input          []string
	expectedResult int
}

type charCountsTest struct {
	input          string
	expectedResult map[int]string
}

type boxIDsTest struct {
	input          []string
	expectedResult string
}

type closeStringTest struct {
	input          string
	input2         string
	expectedResult bool
}

type removeDifferingCharTest struct {
	input          string
	input2         string
	expectedResult string
}

func TestChecksum(t *testing.T) {
	var tests = []checksumTest{
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	}

	for _, test := range tests {
		result := checksum(test.input)
		if result != test.expectedResult {
			t.Fatalf("Checksum returned wrong result, expecting %d got %d", test.expectedResult, result)
		}
	}
}

func TestCharCounts(t *testing.T) {
	var tests = []charCountsTest{
		{"abcdef", map[int]string{1: ""}},
		{"bababc", map[int]string{3: "", 2: "", 1: ""}},
		{"abbcde", map[int]string{2: "", 1: ""}},
		{"abcccd", map[int]string{3: "", 1: ""}},
		{"aabcdd", map[int]string{2: "", 1: ""}},
		{"abcdee", map[int]string{2: "", 1: ""}},
		{"ababab", map[int]string{3: ""}},
	}

	for _, test := range tests {
		result := charCounts(test.input)
		if !reflect.DeepEqual(result, test.expectedResult) {
			t.Fatalf("Char Counts returned wrong result, expecting %v got %v", test.expectedResult, result)
		}
	}

}

func TestBoxIDs(t *testing.T) {
	var tests = []boxIDsTest{
		{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, "fgij"},
		{[]string{"abcdp", "abcde", "fghija", "klmno", "pqrst", "fguijb", "axcye", "wvxyz"}, "abcd"},
	}

	for _, test := range tests {
		result := boxIDs(test.input)
		if result != test.expectedResult {
			t.Fatalf("boxIds returned wrong result, expecting %s got %s", test.expectedResult, result)
		}
	}
}

func TestCloseString(t *testing.T) {
	var tests = []closeStringTest{
		{"abcde", "axcye", false},
		{"fghij", "fguij", true},
	}

	for _, test := range tests {
		result := closeString(test.input, test.input2)
		if result != test.expectedResult {
			t.Fatalf("closeString returned wrong result, expecting %v got %v", test.expectedResult, result)
		}
	}
}

func TestRemoveDifferingChar(t *testing.T) {
	var tests = []removeDifferingCharTest{
		{"fghij", "fguij", "fgij"},
	}

	for _, test := range tests {
		result := removeDifferingChar(test.input, test.input2)
		if result != test.expectedResult {
			t.Fatalf("removeDifferingChar returned wrong result, expecting %s got %s", test.expectedResult, result)
		}
	}
}
