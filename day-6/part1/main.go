package main

import (
	"bufio"
	_ "embed"
	"fmt"

	//    "regexp"
	//    "strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseLine(line string) ([]int, []string) {
	numbers := []int{}

	fields := strings.Fields(line)
	for _, field := range fields {
		if field == "*" || field == "+" {
			return numbers, fields
		}

		var n int
		fmt.Sscanf(field, "%d", &n)
		numbers = append(numbers, n)
	}

	return numbers, []string{}
}

func parseLines(i string) ([][]int, []string) {
	numbers := [][]int{}
	operations := []string{}

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()

		var numbersLine []int
		numbersLine, operations = parseLine(line)
		if len(numbersLine) == 0 {
			continue
		}

		numbers = append(numbers, numbersLine)
	}

	return numbers, operations
}

func reverseIntMatrices(input [][]int) [][]int {
	if len(input) == 0 {
		return input
	}

	numRows := len(input)
	numCols := len(input[0])

	output := make([][]int, numCols)
	for i := range output {
		output[i] = make([]int, numRows)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			output[j][i] = input[i][j]
		}
	}

	return output
}

func processLines(numbers [][]int, operations []string) int {
	var grandTotal int

	for i, number := range reverseIntMatrices(numbers) {
		var lineTotal int
		for j, n := range number {
			switch operations[i] {
			case "+":
				lineTotal += n
			case "*":
				if j == 0 {
					lineTotal = n
				} else {
					lineTotal *= n
				}
			}
		}

		grandTotal += lineTotal
	}

	return grandTotal
}

func run(i string) int {
	numbers, operations := parseLines(i)

	answer := processLines(numbers, operations)

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
