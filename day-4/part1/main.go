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

func parseLine(line string) []int {
	gridLine := []int{}

	for _, pos := range line {
		if pos == '@' {
			gridLine = append(gridLine, 1)
		} else {
			gridLine = append(gridLine, 0)
		}
	}

	return gridLine
}

// grid is [y][x] ordonned
func parseLines(i string) [][]int {
	grid := [][]int{}
	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		gridLine := parseLine(line)
		grid = append(grid, gridLine)
	}

	return grid
}

// grid is [y][x] ordonned
func sumAdjacentCells(grid [][]int, x, y int) int {
	rolls := 0

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i < 0 || j < 0 || j >= len(grid) || i >= len(grid[0]) {
				continue
			}
			if i == x && j == y {
				continue
			}
			rolls += grid[j][i]
		}
	}

	return rolls
}

// grid is [y][x] ordonned
func getPaperRolls(grid [][]int) int {
	rolls := 0

	for y, line := range grid {
		for x, cell := range line {
			// no roll paper at this position
			if cell == 0 {
				continue
			}

			n := sumAdjacentCells(grid, x, y)
			if n < 4 {
				rolls++
			}
		}
	}

	return rolls
}

func run(i string) int {
	var answer int

	grid := parseLines(i)
	fmt.Println("Grid:", grid)

	answer = getPaperRolls(grid)

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
