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

	// Example bank:
	// 987654328901234
	// len bank = 15
	// Find highest digit from index i to len-11 (len - (12 - round))
	// i = 0: pick highest number from index 0 to 3 -> pick 9 at index 0
	//
	fistIndex := 0
	fmt.Println("Bank for max joltage:", bank)
	for i := 0; i < 12; i++ {
		lastIndex := 0
		remainingDigits := 12 - i
		maxIndex := len(bank) - remainingDigits + 1
		highestDigit := -1
		subBank := bank[fistIndex+i : maxIndex]
		for i, v := range subBank {
			if v > highestDigit {
				highestDigit = v
				lastIndex = i
			}
		}
		fistIndex = fistIndex + lastIndex
		maxJoltage = maxJoltage*10 + highestDigit
		fmt.Printf("\t highestDigit %d: %d at index %d from %v (max Joltage = %d)\n", i, highestDigit, fistIndex, subBank, maxJoltage)
	}

	return maxJoltage
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
