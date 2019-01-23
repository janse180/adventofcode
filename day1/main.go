package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const filename = "input.txt"

func main() {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		integer, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("Error parsing line into int: %v %s\n", err, scanner.Text())
		}
		input = append(input, integer)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	resultPart1 := calibrate(0, input)
	resultPart2 := calibratePart2(0, input)

	fmt.Printf("Result of calibration for part 1: %d\n", resultPart1)
	fmt.Printf("Result of calibration for part 2: %d\n", resultPart2)
}

func calibrate(start int, input []int) int {
	result := start

	for _, i := range input {
		result += i
	}

	return result
}

func calibratePart2(start int, input []int) int {
	result := start

	var results = []int{start}

	for {
		for _, i := range input {
			result += i
			if contains(result, results) {
				return result
			}
			results = append(results, result)
		}
	}
}

func contains(i int, list []int) bool {
	for _, item := range list {
		if i == item {
			return true
		}
	}
	return false
}
