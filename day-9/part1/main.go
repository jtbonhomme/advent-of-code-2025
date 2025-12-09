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

type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// 12,1
func parseLine(line string) Position {
	var p Position
	parts := strings.Split(line, ",")
	fmt.Sscanf(parts[0], "%d", &p.X)
	fmt.Sscanf(parts[1], "%d", &p.Y)
	return p
}

func parseLines(i string) []Position {
	res := []Position{}
	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, parseLine(line))
	}

	return res
}

func processLines(lines []Position) int {
	var maxArea int = -1

	// parse all position couples
	// and find the biggest area
	for _, p1 := range lines {
		area := 0
		for _, p2 := range lines {
			if p1 == p2 {
				continue
			}
			area = (p2.X - p1.X + 1) * (p2.Y - p1.Y + 1)
			if area < 0 {
				area = -area
			}
			//fmt.Printf("Area between %s and %s = %d\n", p1, p2, area)
			// find maximum area
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
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
