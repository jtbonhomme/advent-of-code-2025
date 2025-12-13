package main

import (
	"bufio"
	"cmp"
	_ "embed"
	"fmt"

	//    "regexp"
	//    "strconv"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

var test bool

func debug(format string, a ...any) {
	if test {
		fmt.Printf(format, a...)
	}
}

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

func getMaxX(positions []Position) int {
	maxX := positions[0].X
	for _, p := range positions {
		if p.X > maxX {
			maxX = p.X
		}
	}
	return maxX
}

func getMaxY(positions []Position) int {
	maxY := positions[0].Y
	for _, p := range positions {
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxY
}

func displayBoard(positions []Position) {
	totalCorners := 0

	debug("Displaying board:\n")
	for i, pos := range positions {
		debug("%02d: %v, ", i+1, pos)
	}
	debug("\n")
	debug("\n   ")
	for x := 0; x <= getMaxX(positions); x++ {
		debug("% 2d ", x)
	}
	debug("\n")

	for y := 0; y <= getMaxY(positions); y++ {
		debug("%02d ", y)
		for x := 0; x <= getMaxX(positions); x++ {
			if slices.Contains(positions, Position{X: x, Y: y}) {
				// found a corner
				totalCorners++
				debug(" %d ", totalCorners)

				continue
			}

			debug(" . ")
		}

		debug("\n")
	}
}

func displayBoardWithRectangle(positions []Position, p1, p2 Position) {
	totalCorners := 0

	debug("Displaying board:\n")
	for i, pos := range positions {
		debug("%02d: %v, ", i+1, pos)
	}
	debug("\n")
	debug("\n   ")
	for x := 0; x <= getMaxX(positions); x++ {
		debug("% 2d ", x)
	}
	debug("\n")

	for y := 0; y <= getMaxY(positions); y++ {
		debug("%02d ", y)
		for x := 0; x <= getMaxX(positions); x++ {
			if slices.Contains(positions, Position{X: x, Y: y}) {
				// found a corner
				totalCorners++
				if (x == p1.X && y == p1.Y) || (x == p2.X && y == p2.Y) {
					// vertical edge of rectangle
					debug("[#]")
				} else {
					debug(" %d ", totalCorners)
				}

				continue
			}

			debug(" . ")
		}

		debug("\n")
	}
}

func processLines(positions []Position) int {
	var maxArea int = -1

	// sort position by ascending Y then X
	slices.SortFunc(positions,
		func(a, b Position) int {
			if a.Y != b.Y {
				return cmp.Compare(a.Y, b.Y)
			}
			return cmp.Compare(a.X, b.X)
		})

	// collect point pairs per row
	var tilesPerRow = make(map[int][]Position)
	for _, p := range positions {
		rowPositions, ok := tilesPerRow[p.Y]
		if !ok {
			// if we did not have any position for this row yet
			rowPositions = []Position{}
		}
		rowPositions = append(rowPositions, p)
		tilesPerRow[p.Y] = rowPositions
	}

	nRows := len(tilesPerRow)

	debug("\nNumber of rows: %d\n", nRows)

	minX := getMaxX(positions) + 1
	maxX := -1
	debug("board dimension: %d,%d\n", getMaxX(positions), getMaxY(positions))
	debug("Min X: %d, Max X: %d\n", minX, maxX)
	fmt.Printf("ordered positions: %v\n", positions)

	// parse all rows and cols occupied ranges
	rowsRanges := computeRowsRanges(positions)
	displayBoard(positions)

	fmt.Printf("\n\n >> Now find all rectangles that can be formed within the new positions\n\n")
	ops := 0
	for i, p1 := range positions {
		area := 0
		for j, p2 := range positions {
			if j >= i {
				continue
			}
			debug("testing positions %v and %v\n", p1, p2)

			ops++
			// We pick two positions that represents the opposite corners of a rectangle
			otherCorner1 := Position{X: p1.X, Y: p2.Y}
			otherCorner2 := Position{X: p2.X, Y: p1.Y}

			// we need to make sure:
			// A) the two other corners are also in the shape
			if !isInTheShape(otherCorner1, rowsRanges) {
				debug("  other corner %v is NOT in the shape\n", otherCorner1)
				continue
			}
			if !isInTheShape(otherCorner2, rowsRanges) {
				debug("  other corner %v is NOT in the shape\n", otherCorner2)
				continue
			}
			// B) each edge of the rectangle is included in the shape
			// C) the center of the rectangle is also in the shape
			center := Position{X: (p1.X + p2.X) / 2, Y: (p1.Y + p2.Y) / 2}
			if !isInTheShape(center, rowsRanges) {
				debug("  center %v is NOT in the shape\n", center)
				continue
			}

			displayBoardWithRectangle(positions, p1, p2)

			// this rectangle is valid, compute area and compare to max area
			area = (abs(p2.X-p1.X) + 1) * (abs(p2.Y-p1.Y) + 1)
			if area < 0 {
				area = -area
			}
			debug("  Found rectangle corners %v and %v with area %d\n\n", p1, p2, area)
			// find maximum area
			if area > maxArea {
				maxArea = area
			}
		}
	}

	debug("Total operations: %d\n", ops)
	return maxArea
}

func computeRowsRanges(positions []Position) map[int][]int {
	rowsRanges := make(map[int][]int) // Y -> [minX, maxX]
	// get first couple of positions (superior horizontal shape edge)
	lastP1 := positions[0]
	lastP2 := positions[1]
	rowsRanges[lastP1.Y] = []int{lastP1.X, lastP2.X}

	for i := 1; i < len(positions)/2; i++ {
		p1 := positions[i*2]
		p2 := positions[i*2+1]
		rowsRanges[p1.Y] = []int{p1.X, p2.X}
		lastP1 = p1
		lastP2 = p2
	}
	debug("rows ranges: %v\n", rowsRanges)

	return rowsRanges
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

// isInTheShape implements an horizontal winding: check if point is within the shape defined by rowsRanges.
// It starts from the given position, then examin each position to the up (until begining of the column)
// each time we cross an horizontal edge, we increment a counter.
// When we reach the upper edge of the board, if the counter is odd,
// we are within the shape, if even, we are outside the shape.
func isInTheShape(p Position, rowsRanges map[int][]int) bool {
	countCrossedEdges := 0

	rowRange, ok := rowsRanges[p.Y]
	if ok && (p.X == rowRange[0] || p.X == rowRange[1]) {
		// p is equal to an existing corner
		return true
	}

	for y := p.Y; y >= 0; y-- {
		rowRange, ok := rowsRanges[y]
		if !ok {
			// no edge on this row
			continue
		}
		// sort x from rowRange
		slices.Sort(rowRange)

		if p.X >= rowRange[0] && p.X <= rowRange[1] {
			// we crossed an horizontal edge
			countCrossedEdges++
		}
	}

	// if we crossed an odd number of edges, we are within the shape
	return (countCrossedEdges%2 == 1)
}

func run(i string) int {
	positions := parseLines(i)
	answer := processLines(positions)
	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
