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

func parseLine(line string) [][]int {
	// Example line: "11-22,55-99"
	// Split around the comma
	parts := strings.Split(line, ",")
	var lineData [][]int
	for _, part := range parts {
		// Split around the hyphen
		bounds := strings.Split(part, "-")
		var boundInts []int
		for _, b := range bounds {
			// Convert to int
			var val int
			fmt.Sscanf(b, "%d", &val)
			boundInts = append(boundInts, val)
		}
		lineData = append(lineData, boundInts)
	}

	return lineData
}

// return true if the ID is invalid
func isInvalidID(id int) bool {
	number := fmt.Sprintf("%d", id)
	len := len(number)
	halfLen := len / 2

	// find repead motifs in the number
	// 112112112 (len = 9)
	// 1x9 = 11111111 compared 12112112 -> false
	// 11 x n can not reach full lenght 9%2 != 0 -> false
	// 112 X 3 = 112112112 compared to 112112112 -> true
	// 1121 X n  can not reach full lenght 9%4 != 0 -> false
	// stop here
	for i := 0; i < halfLen; i++ {
		if len%(i+1) != 0 {
			continue
		}
		motif := number[0 : i+1]
		repeated := ""
		times := len / (i + 1)
		for j := 0; j < times; j++ {
			repeated += motif
		}
		if repeated == number {
			fmt.Printf("[%d =  %s x %d] ", id, motif, times)
			return true
		}
	}

	return false
}

func countInvalidIDsInRange(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if isInvalidID(i) {
			count += i
		}
	}

	return count
}

func run(i string) int {
	var answer int

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		res := parseLine(line)

		for _, r := range res {
			fmt.Println("Range:", r)

			answer += countInvalidIDsInRange(r[0], r[1])

			fmt.Println("")
		}
	}

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
