package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filename = "input.txt"

func main() {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	g := newGridCombined()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c, err := parseClaim(scanner.Text())
		if err != nil {
			panic(fmt.Errorf("error parsing claim %s : %v", scanner.Text(), err))
		}
		g.claims = append(g.claims, c)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g.plotClaims()
	result := g.grid.calcOverlapArea()
	resultPart2 := g.getNonOverlapedID()

	fmt.Printf("Overlapping Area %d\n", result)
	fmt.Printf("Result of NonOverlapping ID for part 2: %d\n", resultPart2)
}

type gridCombined struct {
	claims []claim
	grid   grid
}

func newGridCombined() gridCombined {
	return gridCombined{grid: newGrid()}
}

func (g *gridCombined) plotClaims() {
	for _, claim := range g.claims {
		g.grid.plot(claim)
	}
}

// Walk over each claim and see if we overlap
func (g *gridCombined) getNonOverlapedID() int {

	for _, c := range g.claims {
		overlap := false
		startX := c.leftEdge
		startY := c.topEdge
		// Y axis
		for y := startY; y <= c.height+c.topEdge-1; y++ {
			// X axis
			for x := startX; x <= c.width+c.leftEdge-1; x++ {
				if len(g.grid[y][x]) > 1 || len(g.grid[y][x]) == 0 {
					overlap = true
				}
			}
		}

		if !overlap {
			return c.id
		}
	}

	//No claim found
	return -1
}

// Y, X, Array of Claim Ids
type grid map[int]map[int][]int

func newGrid() grid {
	return make(grid)
}

// Plots the claim on the grid
func (g *grid) plot(c claim) {
	startX := c.leftEdge
	startY := c.topEdge

	// Y axis
	for y := startY; y <= c.height+c.topEdge-1; y++ {
		// X axis
		for x := startX; x <= c.width+c.leftEdge-1; x++ {
			if (*g)[y] == nil {
				(*g)[y] = map[int][]int{}
			}
			(*g)[y][x] = append((*g)[y][x], c.id)
		}
	}
}

func (g *grid) calcOverlapArea() int {
	total := 0

	for _, row := range *g {
		for _, val := range row {
			if len(val) > 1 {
				total++
			}
		}
	}

	return total
}

func (g *grid) print() {
	var yKeys []int
	var xKeys = map[int][]int{}
	for ky, row := range *g {
		yKeys = append(yKeys, ky)
		for kx := range row {
			xKeys[ky] = append(xKeys[ky], kx)
		}
		sort.Ints(xKeys[ky])
	}
	sort.Ints(yKeys)
	// Header
	for _, y := range yKeys {
		for _, x := range xKeys[y] {
			fmt.Printf("%v", (*g)[y][x])
		}
		fmt.Printf(" %v \n", y)
	}
}

type claim struct {
	id       int
	leftEdge int
	topEdge  int
	width    int
	height   int
}

func parseClaim(in string) (claim, error) {
	c := claim{}
	splitIn := strings.Split(in, " ")

	id, err := strconv.Atoi(strings.Trim(splitIn[0], "#"))
	if err != nil {
		return c, err
	}
	c.id = id

	edges := strings.Split(strings.Trim(splitIn[2], ":"), ",")
	leftEdge, err := strconv.Atoi(edges[0])
	if err != nil {
		return c, err
	}
	c.leftEdge = leftEdge

	topEdge, err := strconv.Atoi(edges[1])
	if err != nil {
		return c, err
	}
	c.topEdge = topEdge

	wh := strings.Split(splitIn[3], "x")
	width, err := strconv.Atoi(wh[0])
	if err != nil {
		return c, err
	}
	c.width = width

	height, err := strconv.Atoi(wh[1])
	if err != nil {
		return c, err
	}
	c.height = height

	return c, nil
}
