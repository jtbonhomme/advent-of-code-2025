package main

import (
	"bufio"
	_ "embed"
	"fmt"

	//    "regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseRangeLine(line string) []int {
	rangeIDs := []int{}
	// split around '-'
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return rangeIDs
	}
	startID, _ := strconv.Atoi(parts[0])
	endID, _ := strconv.Atoi(parts[1])
	for id := startID; id <= endID; id++ {
		rangeIDs = append(rangeIDs, id)
	}
	return rangeIDs
}

func parseProductIDLine(line string) int {
	var id int
	id, _ = strconv.Atoi(line)
	return id
}

// parse input tests
func parseLines(i string) ([]int, []int) {
	productIDs := []int{}
	freshIDs := []int{}

	scanner := bufio.NewScanner(strings.NewReader(i))

	var foundBlankLine bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			foundBlankLine = true
			continue
		}

		if foundBlankLine {
			id := parseProductIDLine(line)
			productIDs = append(productIDs, id)
		} else {
			rangeIDs := parseRangeLine(line)
			freshIDs = append(freshIDs, rangeIDs...)
		}
	}

	return freshIDs, productIDs
}

// returns number of fresh products
func processLines(freshIDs []int, availableProductIDs []int) int {
	var result int

	for _, id := range availableProductIDs {
		for _, freshID := range freshIDs {
			if id == freshID {
				result++
				break
			}
		}
	}

	return result
}

func run(i string) int {
	freshIDs, availableProductIDs := parseLines(i)

	answer := processLines(freshIDs, availableProductIDs)

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
