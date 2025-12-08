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

func parseLine(line string, input map[int]int) map[int]int {
	var output = make(map[int]int)

	// display input line
	/*	for i := 0; i < len(line); i++ {
			_, ok := input[i]
			if ok {
				fmt.Printf("%d", input[i])
			} else {
				fmt.Printf(".")
			}
		}

		fmt.Println()
		fmt.Println(line)*/

	for i, c := range line {
		switch c {
		case '.':
			_, ok := input[i]
			if ok {
				output[i] += input[i]
			}
			continue
		case 'S':
			// entry point
			output[i] = 1
		case '^':
			// splitter
			_, ok := input[i]
			if !ok {
				continue
			}

			if i > 0 {
				output[i-1] += input[i]
			}
			if i < len(line)-1 {
				output[i+1] += input[i]
			}
		}
	}

	// display line
	for i := 0; i < len(line); i++ {
		_, ok := input[i]
		if ok {
			fmt.Printf("%d", input[i])
		} else {
			fmt.Printf(".")
		}
	}

	fmt.Println()

	return output
}

func parseLines(i string) map[int]int {
	res := make(map[int]int)

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		res = parseLine(line, res)
	}

	return res
}

func processLines(beams map[int]int) int {
	res := 0
	for _, v := range beams {
		res += v
	}

	return res
}

func run(i string) int {
	lines := parseLines(i)

	answer := processLines(lines)

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
