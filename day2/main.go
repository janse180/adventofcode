package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const filename = "input.txt"

func main() {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := checksum(input)
	// resultPart2 := calibratePart2(0, input)
	//
	fmt.Printf("Result of checksum %d\n", result)
	// fmt.Printf("Result of calibration for part 2: %d\n", resultPart2)
}

func checksum(input []string) int {

	countTwo := 0
	countThree := 0

	for _, i := range input {
		counts := charCounts(i)
		if _, ok := counts[2]; ok {
			countTwo++
		}
		if _, ok := counts[3]; ok {
			countThree++
		}
	}

	return countTwo * countThree
}

func charCounts(input string) map[int]string {

	result := map[string]int{}
	resultNoDuplicates := map[int]string{}

	for _, c := range input {
		result[string(c)]++
	}

	// Don't return duplicate counts
	for _, v := range result {
		resultNoDuplicates[v] = ""
	}

	return resultNoDuplicates
}
