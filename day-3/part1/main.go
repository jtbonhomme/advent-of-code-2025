package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"

	//    "regexp"
	//    "strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseLine(line string) []int {
	// Example line: "1122"
	lineData := []int{}
	for _, digit := range line {
		i, _ := strconv.Atoi(string(digit))
		lineData = append(lineData, i)
	}

	return lineData
}

func getMaxJoltage(bank []int) int {
	maxJoltage := 0

	if len(bank) == 0 {
		return maxJoltage
	}

	// bank: 218
	// 21 > 18 > 8
	// Look for first highest number in the bank range, except last digit
	// Then look for the second highest number in the bank range, except first highest digit
	firstHighest := -1
	firstHighestIndex := -1
	secondHighest := -1
	for i, v := range bank {
		if i == len(bank)-1 {
			break
		}
		if v > firstHighest {
			firstHighest = v
			firstHighestIndex = i
		}
	}

	for i, v := range bank {
		if i <= firstHighestIndex {
			continue
		}
		if v > secondHighest {
			secondHighest = v
		}
	}

	return firstHighest*10 + secondHighest
}

func run(i string) int {
	var answer int

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		res := parseLine(line)

		fmt.Println("Bank:", res)

		answer += getMaxJoltage(res)

		fmt.Println("")
	}

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
