package main

import (
	"bufio"
	"cmp"
	"embed"
	_ "embed"
	"flag"
	"fmt"
	"log"

	//    "regexp"
	//    "strconv"
	"slices"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/schollz/progressbar/v3"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

type Game struct {
}

var isStarted bool
var pixels []byte

//go:embed *.txt
var inputs embed.FS
var inputfile string

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

func getMinX(positions []Position) int {
	if len(positions) == 0 {
		return -1
	}

	minX := positions[0].X
	for _, p := range positions {
		if p.X < minX {
			minX = p.X
		}
	}
	return minX
}

func getMinY(positions []Position) int {
	if len(positions) == 0 {
		return -1
	}

	minY := positions[0].Y
	for _, p := range positions {
		if p.Y < minY {
			minY = p.Y
		}
	}
	return minY
}

func getMaxX(positions []Position) int {
	if len(positions) == 0 {
		return -1
	}

	maxX := positions[0].X
	for _, p := range positions {
		if p.X > maxX {
			maxX = p.X
		}
	}
	return maxX
}

func getMaxY(positions []Position) int {
	if len(positions) == 0 {
		return -1
	}

	maxY := positions[0].Y
	for _, p := range positions {
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxY
}

func displayBoard(positions []Position) {
	if !test {
		return
	}

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
	if !test {
		return
	}

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
	debug("ordered positions: %v\n", positions)

	// parse all rows and cols occupied ranges
	fmt.Println("computeRowsRanges")
	rowsRanges := computeRowsRanges(positions)
	fmt.Printf("found %d rows ranges\n", len(rowsRanges))
	displayBoard(positions)

	fmt.Println("find all rectangles that can be formed within the new positions")
	ops := 0
	bar := progressbar.Default(int64(firstIntegerSum(len(positions))))
	for i, p1 := range positions {
		area := 0
		for j, p2 := range positions {
			if j >= i {
				continue
			}
			debug("testing positions: %v and %v\n", p1, p2)
			bar.Add(1)
			ops++
			// We pick two positions that represents the opposite corners of a rectangle
			otherCorner1 := Position{X: p1.X, Y: p2.Y}
			otherCorner2 := Position{X: p2.X, Y: p1.Y}
			debug("other corners: %v and %v\n", otherCorner1, otherCorner2)

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

	bar := progressbar.Default(int64(len(positions)/2 - 1))

	for i := 1; i < len(positions)/2; i++ {
		bar.Add(1)
		p1 := positions[i*2]
		p2 := positions[i*2+1]
		rowsRanges[p1.Y] = []int{p1.X, p2.X}
		// sort x from rowRange
		slices.Sort(rowsRanges[p1.Y])
		lastP1 = p1
		lastP2 = p2
	}

	return rowsRanges
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func firstIntegerSum(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		res += i
	}

	return res
}

// isInTheShape implements an horizontal winding: check if point is within the shape defined by rowsRanges.
// It starts from the given position, then examin each position to the up (until begining of the column)
// each time we cross an horizontal edge, we increment a counter.
// When we reach the upper edge of the board, if the counter is odd,
// we are within the shape, if even, we are outside the shape.
func isInTheShape(p Position, rowsRanges map[int][]int) bool {
	countCrossedEdges := 0
	// first check if p is exactly on an edge
	rowRange, ok := rowsRanges[p.Y]
	if ok &&
		p.X >= rowRange[0] &&
		p.X <= rowRange[1] {
		// p is located on an horizontal edge
		return true
	}

	// the point is not on an edge, we can proceed with the winding
	for y := p.Y; y >= 0; y-- {
		rowRange, ok := rowsRanges[y]
		if !ok {
			// no edge on this row
			continue
		}

		if p.X >= rowRange[0] && p.X <= rowRange[1] {
			// we crossed an horizontal edge
			countCrossedEdges++
		}

		// if we reached the bottom of a vertical edge,
		// we move y to the upper vertical corner
		if p.X == rowRange[0] || p.X == rowRange[1] {
			// is there another corner at the same column?
			nextY := -1
			for yy := y - 1; yy >= 0; yy-- {
				rr, ok := rowsRanges[yy]
				if !ok {
					continue
				}
				if rr[0] == p.X || rr[1] == p.X {
					nextY = yy
					break
				}
			}
			if nextY != -1 {
				y = nextY
			}
		}
	}

	// if we crossed an odd number of edges, we are within the shape
	rest := countCrossedEdges % 2
	return (rest == 1)
}

var positions []Position

func getScaleFactor() float64 {

	minX := getMinX(positions)
	minY := getMinY(positions)
	maxX := getMaxX(positions)
	maxY := getMaxY(positions)

	scaleX := 1.0
	if maxX-minX >= screenWidth {
		scaleX = float64(screenWidth) / float64(maxX-minX+1)
	}
	scaleY := 1.0
	if maxY-minY >= screenHeight {
		scaleY = float64(screenHeight) / float64(maxY-minY+1)
	}

	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	return scale
}

func draw(pixels []byte) {
	if len(pixels) != screenWidth*screenHeight*4 {
		log.Fatalf("pixels length is %d, expected %d", len(pixels), screenWidth*screenHeight*4)
	}

	// clear screen
	for i := range pixels {
		pixels[i] = 0x00
	}

	if len(positions) == 0 {
		return
	}

	scale := getScaleFactor()

	// draw positions
	for _, p := range positions {
		scaledX := int(float64(p.X) * scale)
		scaledY := int(float64(p.Y) * scale)
		p := Position{X: scaledX, Y: scaledY}
		if p.X < 0 || p.X >= screenWidth || p.Y < 0 || p.Y >= screenHeight {
			continue
		}
		idx := (p.Y*screenWidth + p.X) * 4
		pixels[idx] = 0xff   // R
		pixels[idx+1] = 0xff // G
		pixels[idx+2] = 0xff // B
		pixels[idx+3] = 0xff // A
	}
}

func run(i string) int {
	fmt.Println("parsing input...")
	positions = parseLines(i)
	fmt.Println("processing input...")
	answer := processLines(positions)
	fmt.Println("... done")
	return answer
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game ended by user")
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("starting processing...")
		if !isStarted {
			isStarted = true
			go func() {
				data, _ := inputs.ReadFile(inputfile)
				answer := run(string(data))
				fmt.Println("Answer: ", answer)
			}()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !isStarted {
		return
	}
	if pixels == nil {
		pixels = make([]byte, screenWidth*screenHeight*4)
	}

	draw(pixels)

	screen.WritePixels(pixels)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("minX: %d - maxX: %d\nminY: %d - maxY: %d\nscale factor: %0.6f", getMinX(positions), getMaxX(positions), getMinY(positions), getMaxY(positions), getScaleFactor()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	flag.StringVar(&inputfile, "input", "input.txt", "input file")
	flag.Parse()
	g := &Game{}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Advent Of Code 2025 - day9 part2")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
