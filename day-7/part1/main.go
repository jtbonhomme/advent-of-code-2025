package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"

	//    "regexp"
	//    "strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseLine(line string, input []int) ([]int, int) {
	var output = make([]int, len(input))
	_ = copy(output, input)
	var split int

	for i, c := range line {
		switch c {
		case '.':
			continue
		case 'S':
			// entry point
			output = []int{i}
		case '^':
			// splitter
			if !slices.Contains(input, i) {
				continue
			}

			// Remove the current index from output
			index := slices.Index(output, i)
			output = slices.Delete(output, index, index+1)

			if i > 0 && !slices.Contains(output, i-1) {
				output = append(output, i-1)
			}
			if i < len(line)-1 && !slices.Contains(output, i+1) {
				output = append(output, i+1)
			}

			split++
		}
	}

	// display line
	for i := 0; i < len(line); i++ {
		if slices.Contains(output, i) && line[i] == '.' {
			fmt.Print("|")
		} else {
			fmt.Printf("%c", line[i])
		}
	}
	fmt.Println()

	return output, split
}

func parseLines(i string) ([]int, int) {
	res := []int{}
	var split, totalSplit int

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		res, split = parseLine(line, res)
		totalSplit += split
	}

	return res, totalSplit
}

func processLines(beams []int, totalSplit int) int {
	return totalSplit
}

func run(i string) int {
	lines, totalSplit := parseLines(i)

	answer := processLines(lines, totalSplit)

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
