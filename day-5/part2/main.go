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
	// split around '-'
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return []int{}
	}
	startID, _ := strconv.Atoi(parts[0])
	endID, _ := strconv.Atoi(parts[1])

	return []int{startID, endID}
}

// parse input lines
// returns freshIDs
func parseLines(i string) [][]int {
	freshIDs := [][]int{}

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rangeIDs := parseRangeLine(line)
		freshIDs = append(freshIDs, rangeIDs)
	}

	return freshIDs
}

type IDRange struct {
	start int
	end   int
}

type IDRanges struct {
	ranges []IDRange
}

func addRange(existingRanges []IDRange, newRange IDRange) []IDRange {
	result := []IDRange{}

	fmt.Printf("Adding new range %v to existing ranges %v\n", newRange, existingRanges)

	if len(existingRanges) == 0 {
		fmt.Printf("a. adding range %v since it is the first range\n", newRange)
		result = append(result, newRange)
		fmt.Printf("a. result is %v\n", result)

		return result
	}

	// check if newRange overlaps with any existing range
	var overlappingRanges []IDRange
	var notOverlappingRanges []IDRange
	for _, existingRange := range existingRanges {
		if newRange.end < existingRange.start || newRange.start > existingRange.end {
			// no overlap
			notOverlappingRanges = append(notOverlappingRanges, existingRange)
			continue
		}

		overlappingRanges = append(overlappingRanges, existingRange)
	}

	if len(overlappingRanges) == 0 {
		// no overlaps, just add the new range
		result = append(existingRanges, newRange)
		fmt.Printf("b. adding range %v since it does not overlap with other ranges\n", newRange)
		fmt.Printf("b. result is %v\n", result)

		return result
	}

	// merge overlapping ranges
	mergedRange := IDRange{
		start: newRange.start,
		end:   newRange.end,
	}

	for _, overlap := range overlappingRanges {
		if overlap.start < mergedRange.start {
			mergedRange.start = overlap.start
		}
		if overlap.end > mergedRange.end {
			mergedRange.end = overlap.end
		}
	}

	fmt.Printf("c. merged range is %v\n", mergedRange)
	result = append(notOverlappingRanges, mergedRange)

	fmt.Printf("c. result is %v\n", result)

	return result
}

// returns number of fresh products
func processLines(freshIDs [][]int) int {
	var result int
	idRanges := []IDRange{}

	for _, freshID := range freshIDs {
		// create IDRange
		idRange := IDRange{
			start: freshID[0],
			end:   freshID[1],
		}

		idRanges = addRange(idRanges, idRange)
	}

	// count total number of IDs in ranges
	for _, r := range idRanges {
		result += (r.end - r.start + 1)
	}

	return result
}

func run(i string) int {
	freshIDs := parseLines(i)
	answer := processLines(freshIDs)
	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
