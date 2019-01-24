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
	resultPart2 := boxIDs(input)

	fmt.Printf("Result of checksum %d\n", result)
	fmt.Printf("Result of boxIDs for part 2: %s\n", resultPart2)
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

func boxIDs(input []string) string {
	for i, id := range input {
		for _, id2 := range input[i+1:] {
			if closeString(id, id2) {
				return removeDifferingChar(id, id2)
			}
		}
	}
	return ""
}

func removeDifferingChar(s1, s2 string) string {
	var differingCharIndex int
	for i, c := range s1 {
		if string(c) != string(s2[i]) {
			differingCharIndex = i
			break
		}
	}

	return s1[:differingCharIndex] + s1[differingCharIndex+1:]
}

// Returns true if there is only a one charecter difference between the strings
func closeString(s1, s2 string) bool {
	notEqualCount := 0
	for i, c := range s1 {
		if notEqualCount > 1 {
			return false
		}
		if string(c) != string(s2[i]) {
			notEqualCount++
		}
	}

	if notEqualCount == 1 {
		return true
	}
	return false
}
