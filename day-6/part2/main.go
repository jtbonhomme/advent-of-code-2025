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

func splitLines(i string) []string {
	lines := []string{}

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func parseOperations(lines string) ([]string, []int) {
	operations := []string{string(lines[0])}
	numColsPerOperation := []int{}

	var currentRows int = 1
	for _, c := range lines[1:] {
		if c == ' ' {
			currentRows++
			continue
		}
		operations = append(operations, string(c))
		numColsPerOperation = append(numColsPerOperation, currentRows)
		currentRows = 1
	}

	numColsPerOperation = append(numColsPerOperation, currentRows)
	return operations, numColsPerOperation
}

func parseLines(lines []string) ([][]int, []string) {
	numbers := [][]int{}

	operations, numColsPerOperation := parseOperations(lines[len(lines)-1])

	numRows := len(lines) - 1

	accumulatedColIndex := 0
	for opIndex := 0; opIndex < len(operations); opIndex++ {
		opNumbers := []int{}
		for colIndex := 0; colIndex < numColsPerOperation[opIndex]; colIndex++ {
			var accumulatedNumber int
			for rowIndex := 0; rowIndex < numRows; rowIndex++ {
				var n int
				c := lines[rowIndex][accumulatedColIndex]
				fmt.Sscanf(string(c), "%d", &n)
				if n != 0 {
					accumulatedNumber = accumulatedNumber*10 + n
				}
			}

			if accumulatedNumber != 0 {
				opNumbers = append(opNumbers, accumulatedNumber)
			}

			accumulatedColIndex++
		}

		numbers = append(numbers, opNumbers)
	}

	return numbers, operations
}

func processLines(numbers [][]int, operations []string) int {
	var grandTotal int

	for i, number := range numbers {
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
	lines := splitLines(i)

	numbers, operations := parseLines(lines)

	answer := processLines(numbers, operations)

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
