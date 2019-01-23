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
