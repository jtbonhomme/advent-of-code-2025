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

	fmt.Printf("\nNumber of rows: %d\n", nRows)
	//sumRows := 0

	minX := getMaxX(positions) + 1
	maxX := -1
	fmt.Printf("Min X: %d, Max X: %d\n", minX, maxX)

	// parse all position couples
	// and find the biggest area
	newPositions := []Position{}

	totalCorners := 0 // 0 to nRows*2 -1
	lastMinX := 0
	lastMaxX := 0
	for y := 0; y <= getMaxY(positions); y++ {
		lineCorners := 0 // 0 to 1
		for x := 0; x <= getMaxX(positions); x++ {
			if slices.Contains(positions, Position{X: x, Y: y}) {
				// found a corner
				totalCorners++
				lineCorners++
				newPositions = append(newPositions, Position{X: x, Y: y})
				fmt.Printf("%03d", totalCorners)

				if totalCorners == 1 {
					// first line first corner
					minX = x
					continue
				}
				if totalCorners == 2 {
					// first line second corner
					maxX = x
					continue
				}
				if totalCorners == nRows*2-1 {
					// last line first corner
					minX = getMaxX(positions) + 1
					continue
				}
				if totalCorners == nRows*2 {
					// last line second corner
					maxX = -1
					continue
				}

				if lineCorners == 1 && x < minX {
					minX = x
					continue
				}

				if lineCorners == 2 && x > maxX {
					maxX = x
					continue
				}

				if lineCorners == 1 && x > minX {
					tiles := tilesPerRow[y]
					secondtile := tiles[1]
					if secondtile.X > maxX {
						maxX = secondtile.X
					}

					continue
				}

				if lineCorners == 2 && x > minX {
					//tiles := tilesPerRow[y]
					//firsttile := tiles[0]
					if lastMaxX < maxX {
						continue
					}
					if minX == lastMinX {
						minX = x
						continue
					}
				}

				continue
			}

			if totalCorners == 1 && x >= minX {
				newPositions = append(newPositions, Position{X: x, Y: y})
				fmt.Printf(" x ")
				continue
			}

			if totalCorners == nRows*2-1 && x <= maxX {
				newPositions = append(newPositions, Position{X: x, Y: y})
				fmt.Printf(" x ")
				continue
			}

			if x <= maxX && x >= minX {
				newPositions = append(newPositions, Position{X: x, Y: y})
				fmt.Printf(" x ")
				continue
			}

			fmt.Printf(" . ")
		}

		fmt.Printf("   minX: %02d, maxX: %02d - lastMinX: %02d, lastMaxX: %02d\n", minX, maxX, lastMinX, lastMaxX)
		lastMinX = minX
		lastMaxX = maxX
	}

	for _, p1 := range newPositions {
		area := 0
		for _, p2 := range positions {
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
	positions := parseLines(i)
	answer := processLines(positions)
	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
