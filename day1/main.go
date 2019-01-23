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

	result := calibrate(0, input)

	fmt.Printf("Result of calibration: %d\n", result)
}

func calibrate(start int, input []int) int {
	result := start

	for _, i := range input {
		result += i
	}

	return result
}
