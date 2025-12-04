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

func displayGrid(grid [][]int) {
	fmt.Println()
	for _, line := range grid {
		for _, cell := range line {
			if cell == 1 {
				fmt.Print("@")
			}
			if cell == 0 {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

// grid is [y][x] ordonned
func getPaperRolls(grid [][]int) int {
	rolls := 0

	displayGrid(grid)
	cellsToClean := make([][]bool, len(grid))
	for i := range cellsToClean {
		cellsToClean[i] = make([]bool, len(grid[0]))
	}

	for y, line := range grid {
		for x, cell := range line {
			// no roll paper at this position
			if cell == 0 {
				continue
			}

			n := sumAdjacentCells(grid, x, y)
			if n < 4 {
				rolls++
				cellsToClean[y][x] = true
			}
		}
	}

	for y, _ := range grid {
		for x, _ := range grid[y] {
			if cellsToClean[y][x] {
				grid[y][x] = 0
			}
		}
	}

	return rolls
}

// grid is [y][x] ordonned
func getAllPaperRolls(grid [][]int) int {
	var answer int

	var stop bool
	for !stop {
		rolls := getPaperRolls(grid)
		fmt.Printf("Rolls this round: %d\n", rolls)
		if rolls == 0 {
			stop = true
		} else {
			answer += rolls
		}
	}

	return answer
}

func run(i string) int {
	var answer int

	grid := parseLines(i)

	answer = getAllPaperRolls(grid)

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
