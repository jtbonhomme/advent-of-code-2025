package main

import (
	"bufio"
	"cmp"
	"embed"
	_ "embed"
	"flag"
	"fmt"
	"image/color"
	"log"
	"time"

	//    "regexp"
	//    "strconv"
	"slices"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth                = 400
	screenHeight               = 400
	processTime  time.Duration = 5 * time.Second
)

type Game struct {
}

var isStarted bool
var pixels []byte

//go:embed *.txt
var inputs embed.FS
var inputfile string

var test bool
var info bool

var rowsRanges map[int][]int
var colsRanges map[int][]int

type Rectangle struct {
	X1 float32
	Y1 float32
	X2 float32
	Y2 float32
}

var candidateRectangle Rectangle
var maxAreaRectangle Rectangle

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

func norm(p1, p2 Position) Position {
	var dx, dy int

	if p1.X != p2.X {
		dx = (p2.X - p1.X) / abs(p2.X-p1.X)
	}
	if p1.Y != p2.Y {
		dy = (p2.Y - p1.Y) / abs(p2.Y-p1.Y)
	}

	return Position{X: dx, Y: dy}
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
	rowsRanges = computeRowsRanges(positions)
	colsRanges = computeColsRanges(positions)
	debug("found %d rows ranges: %v\n", len(rowsRanges), rowsRanges)
	debug("found %d cols ranges: %v\n", len(colsRanges), colsRanges)
	//displayBoard(positions)

	fmt.Println("find all rectangles that can be formed within the new positions")
	ops := 0
	for i, p1 := range positions {
		area := 0
		for j, p2 := range positions {
			if j >= i {
				continue
			}
			debug("testing positions: %v and %v\n", p1, p2)
			ops++
			if test {
				time.Sleep(time.Duration(250 * time.Millisecond))
			}

			// We pick two positions that represents the opposite corners of a rectangle
			otherCorner1 := Position{X: p1.X, Y: p2.Y}
			otherCorner2 := Position{X: p2.X, Y: p1.Y}
			debug("other corners: %v and %v\n", otherCorner1, otherCorner2)

			candidateRectangle = Rectangle{
				X1: float32(p1.X),
				Y1: float32(p1.Y),
				X2: float32(p2.X),
				Y2: float32(p2.Y),
			}

			// we need to make sure that:
			// A) the two other corners are also in the shape
			if !isInTheShape(otherCorner1, rowsRanges, colsRanges) {
				debug("  other corner %v is NOT in the shape\n", otherCorner1)
				continue
			}
			if !isInTheShape(otherCorner2, rowsRanges, colsRanges) {
				debug("  other corner %v is NOT in the shape\n", otherCorner2)
				continue
			}

			// B) the interior of the rectangle is included in the shape
			// top edge center
			p1ToP2 := norm(p1, p2)
			testP1 := Position{X: p1.X + p1ToP2.X, Y: p1.Y + p1ToP2.Y}
			if !isInTheShape(testP1, rowsRanges, colsRanges) {
				debug("  top edge center %v is NOT in the shape\n", testP1)
				continue
			}
			// bottom edge center
			p2ToP1 := norm(p2, p1)
			testP2 := Position{X: p2.X + p2ToP1.X, Y: p2.Y + p2ToP1.Y}
			if !isInTheShape(testP2, rowsRanges, colsRanges) {
				debug("  bottom edge center %v is NOT in the shape\n", testP2)
				continue
			}
			// left edge center
			otherCorner1TOotherCorner2 := norm(otherCorner1, otherCorner2)
			testOtherCorner1 := Position{X: otherCorner1.X + otherCorner1TOotherCorner2.X, Y: otherCorner1.Y + otherCorner1TOotherCorner2.Y}
			if !isInTheShape(testOtherCorner1, rowsRanges, colsRanges) {
				debug("  left edge center %v is NOT in the shape\n", testOtherCorner1)
				continue
			}
			// right edge center
			otherCorner2TOotherCorner1 := norm(otherCorner2, otherCorner1)
			testOtherCorner2 := Position{X: otherCorner2.X + otherCorner2TOotherCorner1.X, Y: otherCorner2.Y + otherCorner2TOotherCorner1.Y}
			if !isInTheShape(testOtherCorner2, rowsRanges, colsRanges) {
				debug("  right edge center %v is NOT in the shape\n", testOtherCorner2)
				continue
			}

			// C) the center of the rectangle is also in the shape
			center := Position{X: (p1.X + p2.X) / 2, Y: (p1.Y + p2.Y) / 2}
			if !isInTheShape(center, rowsRanges, colsRanges) {
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
				maxAreaRectangle = candidateRectangle
				debug("  New MAX AREA: %d\n\n", maxArea)
			}
		}
	}

	debug("Total operations: %d\n", ops)
	return maxArea
}

func getBoardDimension(p Position) (int, int) {
	scale := getScaleFactor()
	boardX := int(float64(p.X) / scale)
	boardY := int(float64(p.Y) / scale)

	return boardX, boardY
}

func drawCandidateRectangle(screen *ebiten.Image) {
	// candidate rectangle is defined by 2 opposite corners (X1,Y1) and (X2,Y2)
	// let's compute the two other corners.
	otherP1X := candidateRectangle.X1
	otherP1Y := candidateRectangle.Y2
	otherP2X := candidateRectangle.X2
	otherP2Y := candidateRectangle.Y1

	// where is the top left corner?

	rectX := float32(min(int(candidateRectangle.X1), int(otherP1X), int(otherP2X)))
	rectY := float32(min(int(candidateRectangle.Y1), int(otherP1Y), int(otherP2Y)))
	rectWidth := float32(abs(max(int(candidateRectangle.X1), int(otherP1X), int(otherP2X)) - min(int(candidateRectangle.X1), int(otherP1X), int(otherP2X))))
	rectHeight := float32(abs(max(int(candidateRectangle.Y1), int(otherP1Y), int(otherP2Y)) - min(int(candidateRectangle.Y1), int(otherP1Y), int(otherP2Y))))

	scale := getScaleFactor()
	boardX := float32(float64(rectX) * scale)
	boardY := float32(float64(rectY) * scale)
	width := float32(float64(rectWidth) * scale)
	height := float32(float64(rectHeight) * scale)
	vector.FillRect(screen, boardX, boardY, width, height, color.RGBA{R: 50, G: 50, B: 50, A: 30}, false)
}

func drawBiggestAreaRectangle(screen *ebiten.Image) {
	// candidate rectangle is defined by 2 opposite corners (X1,Y1) and (X2,Y2)
	// let's compute the two other corners.
	otherP1X := maxAreaRectangle.X1
	otherP1Y := maxAreaRectangle.Y2
	otherP2X := maxAreaRectangle.X2
	otherP2Y := maxAreaRectangle.Y1

	// where is the top left corner?

	rectX := float32(min(int(maxAreaRectangle.X1), int(otherP1X), int(otherP2X)))
	rectY := float32(min(int(maxAreaRectangle.Y1), int(otherP1Y), int(otherP2Y)))
	rectWidth := float32(abs(max(int(maxAreaRectangle.X1), int(otherP1X), int(otherP2X)) - min(int(maxAreaRectangle.X1), int(otherP1X), int(otherP2X))))
	rectHeight := float32(abs(max(int(maxAreaRectangle.Y1), int(otherP1Y), int(otherP2Y)) - min(int(maxAreaRectangle.Y1), int(otherP1Y), int(otherP2Y))))

	scale := getScaleFactor()
	boardX := float32(float64(rectX) * scale)
	boardY := float32(float64(rectY) * scale)
	width := float32(float64(rectWidth) * scale)
	height := float32(float64(rectHeight) * scale)
	vector.FillRect(screen, boardX, boardY, width, height, color.RGBA{R: 150, G: 150, B: 150, A: 30}, false)
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
		// sort x from rowRange
		slices.Sort(rowsRanges[p1.Y])
		lastP1 = p1
		lastP2 = p2
	}

	return rowsRanges
}

func computeColsRanges(positions []Position) map[int][]int {
	colsRanges := make(map[int][]int) // X -> [minY, maxY]

	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			if positions[j].X == positions[i].X {
				colsRanges[positions[i].X] = []int{positions[i].Y, positions[j].Y}
			}
		}
	}

	return colsRanges
}

func extendRows(rowsRanges, colsRanges map[int][]int) map[int][]int {

	extended := make(map[int][]int, len(rowsRanges))

	for y, xr := range rowsRanges {
		xmin, xmax := xr[0], xr[1]
		extendedMax := xmax

		for x, yr := range colsRanges {
			ymin, ymax := yr[0], yr[1]

			// la colonne x coupe la ligne y
			if y >= ymin && y <= ymax {
				if x >= xmax && x > extendedMax {
					extendedMax = x
				}
			}
		}

		extended[y] = []int{xmin, extendedMax}
	}

	return extended
}

func extendCols(colsRanges, rowsRanges map[int][]int) map[int][]int {

	extended := make(map[int][]int, len(colsRanges))

	for x, yr := range colsRanges {
		ymin, ymax := yr[0], yr[1]
		extendedMax := ymax

		for y, xr := range rowsRanges {
			xmin, xmax := xr[0], xr[1]

			// la ligne y coupe la colonne x
			if x >= xmin && x <= xmax {
				if y >= ymax && y > extendedMax {
					extendedMax = y
				}
			}
		}

		extended[x] = []int{ymin, extendedMax}
	}

	return extended
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v > m {
			m = v
		}
	}
	return m
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
func isInTheShape(p Position, rowsRanges, colsRanges map[int][]int) bool {
	countCrossedEdges := 0
	// first check if p is exactly on an edge
	extendedRows := extendRows(rowsRanges, colsRanges)
	rowRange, ok := extendedRows[p.Y]
	if ok {
		for i := 0; i < len(rowRange)/2; i++ {
			if p.X >= rowRange[i] &&
				p.X <= rowRange[i+1] {
				// p is located on an horizontal edge
				return true
			}
		}
		return false
	}

	// the point is not on an edge, we can proceed with the winding
	// the point is moved on the x axis from its position to the left of the board
	for x := p.X; x >= 0; x-- {
		colRange, ok := colsRanges[x]
		if !ok {
			// no edge on this row
			continue
		}

		slices.Sort(colRange)
		if p.Y >= colRange[0] && p.Y <= colRange[1] {
			// we crossed an horizontal edge
			countCrossedEdges++
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

	scaleX := float64(screenWidth) / float64(maxX+minX+1)
	scaleY := float64(screenHeight) / float64(maxY+minY+1)

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

	// draw rows ranges
	for y, rowRange := range rowsRanges {
		scaledY := int(float64(y) * scale)
		if scaledY < 0 || scaledY >= screenHeight {
			continue
		}
		for x := int(float64(rowRange[0]) * scale); x <= int(float64(rowRange[1])*scale); x++ {
			if x < 0 || x >= screenWidth {
				continue
			}
			idx := (scaledY*screenWidth + x) * 4
			pixels[idx] = 0x81   // R
			pixels[idx+1] = 0x81 // G
			pixels[idx+2] = 0xf1 // B
			pixels[idx+3] = 0x60 // A
		}
	}

	// draw cols ranges
	for x, colRange := range colsRanges {
		scaledX := int(float64(x) * scale)
		if scaledX < 0 || scaledX >= screenWidth {
			continue
		}
		for y := int(float64(colRange[0]) * scale); y <= int(float64(colRange[1])*scale); y++ {
			if y < 0 || y >= screenHeight {
				continue
			}
			idx := (y*screenWidth + scaledX) * 4
			pixels[idx] = 0x81   // R
			pixels[idx+1] = 0x81 // G
			pixels[idx+2] = 0xf1 // B
			pixels[idx+3] = 0x60 // A
		}
	}

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
		pixels[idx+1] = 0x1f // G
		pixels[idx+2] = 0x1f // B
		pixels[idx+3] = 0x1f // A
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

	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		info = !info
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		boardX, boardY := getBoardDimension(Position{X: x, Y: y})
		fmt.Printf("mouse clicked at pixel (%d,%d) -> board position (%d,%d)\n", x, y, boardX, boardY)
		res := isInTheShape(Position{X: boardX, Y: boardY}, rowsRanges, colsRanges)
		if res {
			fmt.Printf("  position (%d,%d) is WITHIN the shape\n", boardX, boardY)
		} else {
			fmt.Printf("  position (%d,%d) is OUTSIDE the shape\n", boardX, boardY)
		}
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

	// draw background
	draw(pixels)
	screen.WritePixels(pixels)

	// draw biggest area rectangle
	drawBiggestAreaRectangle(screen)

	// draw candidate rectangle
	drawCandidateRectangle(screen)

	// display info
	if info {
		x, y := ebiten.CursorPosition()
		boardX, boardY := getBoardDimension(Position{X: x, Y: y})
		ebitenutil.DebugPrint(screen, fmt.Sprintf("minX: %d - maxX: %d\nminY: %d - maxY: %d\nscale factor: %0.6f\npos: %d, %d | %d, %d", getMinX(positions), getMaxX(positions), getMinY(positions), getMaxY(positions), getScaleFactor(), x, y, boardX, boardY))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	flag.StringVar(&inputfile, "input", "input.txt", "input file")
	flag.BoolVar(&test, "test", false, "input file")
	flag.Parse()
	g := &Game{}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Advent Of Code 2025 - day9 part2")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
