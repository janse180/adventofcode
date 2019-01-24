package main

import (
	"testing"
)

type plotTest struct {
	claim []claim
}

type calcOverlapAreaTest struct {
	claim        []claim
	expectedArea int
}

type parseClaimTest struct {
	input         string
	expectedClaim claim
}

type getNonOverlapIDTest struct {
	input      gridCombined
	expectedID int
}

func TestParseClaim(t *testing.T) {
	var tests = []parseClaimTest{
		{"#123 @ 3,2: 5x4", claim{id: 123, leftEdge: 3, topEdge: 2, width: 5, height: 4}},
	}

	for _, test := range tests {

		result, _ := parseClaim(test.input)
		if result != test.expectedClaim {
			t.Fatalf("parseClaim returned wrong result, expecting %v got %v", test.expectedClaim, result)
		}
	}

}

func TestPlot(t *testing.T) {
	var tests = []plotTest{
		{[]claim{
			{id: 1, leftEdge: 1, topEdge: 3, width: 4, height: 4},
			{id: 2, leftEdge: 3, topEdge: 1, width: 4, height: 4},
			{id: 3, leftEdge: 5, topEdge: 5, width: 2, height: 2},
		}},
	}

	for _, test := range tests {
		g := newGrid()

		for _, claim := range test.claim {
			g.plot(claim)
		}

		// g.print()

		// result := checksum(test.input)
		// if result != test.expectedResult {
		// 	t.Fatalf("Checksum returned wrong result, expecting %d got %d", test.expectedResult, result)
		// }
	}
}

func TestCalcOverlapArea(t *testing.T) {
	var tests = []calcOverlapAreaTest{
		{[]claim{
			{id: 1, leftEdge: 1, topEdge: 3, width: 4, height: 4},
			{id: 2, leftEdge: 3, topEdge: 1, width: 4, height: 4},
			{id: 3, leftEdge: 5, topEdge: 5, width: 2, height: 2},
		}, 4},
	}

	for _, test := range tests {
		g := newGrid()

		for _, claim := range test.claim {
			g.plot(claim)
		}

		result := g.calcOverlapArea()
		if result != test.expectedArea {
			t.Fatalf("calcOverlapArea returned wrong result, expecting %d got %d", test.expectedArea, result)
		}
	}
}

func TestGetNonOverlapId(t *testing.T) {
	var tests = []getNonOverlapIDTest{
		{gridCombined{[]claim{
			{id: 1, leftEdge: 1, topEdge: 3, width: 4, height: 4},
			{id: 2, leftEdge: 3, topEdge: 1, width: 4, height: 4},
			{id: 3, leftEdge: 5, topEdge: 5, width: 2, height: 2},
		}, nil}, 3},
	}

	for _, test := range tests {
		test.input.grid = newGrid()
		test.input.plotClaims()

		result := test.input.getNonOverlapedID()
		if result != test.expectedID {
			t.Fatalf("getNonOverlapedID returned wrong result, expecting %d got %d", test.expectedID, result)
		}
	}
}
