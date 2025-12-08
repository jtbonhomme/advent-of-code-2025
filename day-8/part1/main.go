package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strconv"

	//    "regexp"
	//    "strconv"
	"strings"
)

//go:embed input.txt
var input string

type Position struct {
	X int
	Y int
	Z int
}

func equalPosition(a, b Position) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

func sqrDist(a, b Position) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return int(dx*dx + dy*dy + dz*dz)
}

var _jbid, _circuitid int

type JunctionBox struct {
	ID          int
	CircuitID   int
	Position    Position
	Connections []*JunctionBox
}

func (jb JunctionBox) String() string {
	return fmt.Sprintf("{id: %04d, pos: (%03d,%03d,%03d), circuit: %d}", jb.ID, jb.Position.X, jb.Position.Y, jb.Position.Z, jb.CircuitID)
}

var JunctionBoxes []*JunctionBox

func parseLine(line string) Position {
	coord := strings.Split(line, ",")

	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])
	z, _ := strconv.Atoi(coord[2])

	return Position{X: x, Y: y, Z: z}
}

func parseLines(i string) []*JunctionBox {
	JunctionBoxes = make([]*JunctionBox, 0)

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		pos := parseLine(line)
		JunctionBoxes = append(JunctionBoxes, &JunctionBox{Position: pos, ID: _jbid, CircuitID: -1})
		_jbid++
	}

	return JunctionBoxes
}

func areConnected(jb1, jb2 *JunctionBox) bool {
	for _, conn := range jb1.Connections {
		if conn == jb2 {
			return true
		}
	}

	return false
}

func areInSameCircuit(jb1, jb2 *JunctionBox) bool {
	visited := make(map[*JunctionBox]bool)

	var dfs func(jb *JunctionBox)
	dfs = func(jb *JunctionBox) {
		visited[jb] = true
		for _, conn := range jb.Connections {
			if !visited[conn] {
				dfs(conn)
			}
		}
	}

	// start DFS from the first junction box in the circuit
	dfs(jb1)

	// check if jb2 was visited
	return visited[jb2]
}

func connectTogether(jb1, jb2 *JunctionBox) {
	// connect them
	jb1.Connections = append(jb1.Connections, jb2)
	jb2.Connections = append(jb2.Connections, jb1)

	if jb1.CircuitID == -1 && jb2.CircuitID == -1 {
		jb1.CircuitID = _circuitid
		jb2.CircuitID = _circuitid
		_circuitid++
	} else if jb1.CircuitID != -1 && jb2.CircuitID == -1 {
		jb2.CircuitID = jb1.CircuitID
	} else if jb1.CircuitID == -1 && jb2.CircuitID != -1 {
		jb1.CircuitID = jb2.CircuitID
	} else if jb1.CircuitID != jb2.CircuitID {
		fmt.Printf("... Merging circuits %d and %d ...\n", jb1.CircuitID, jb2.CircuitID)
		oldCircuitID := jb2.CircuitID
		newCircuitID := jb1.CircuitID
		for _, jb := range JunctionBoxes {
			if jb.CircuitID == oldCircuitID {
				jb.CircuitID = newCircuitID
			}
		}
	}
}

func findAllDistances(jbs []*JunctionBox) map[int][]*JunctionBox {
	dists := make(map[int][]*JunctionBox)

	for _, jb := range jbs {
		for _, otherJb := range jbs {
			if equalPosition(jb.Position, otherJb.Position) {
				continue
			}

			dist := sqrDist(jb.Position, otherJb.Position)
			dists[dist] = []*JunctionBox{jb, otherJb}
		}
	}

	return dists
}

func findClosestNonConnectedBoxes(jbs []*JunctionBox) (*JunctionBox, *JunctionBox, int) {
	minDist := -1
	jb1, jb2 := &JunctionBox{}, &JunctionBox{}

	for _, jb := range jbs {
		for _, otherJb := range jbs {
			if equalPosition(jb.Position, otherJb.Position) {
				continue
			}

			if areConnected(jb, otherJb) {
				continue
			}

			dist := sqrDist(jb.Position, otherJb.Position)

			if minDist == -1 || dist < minDist {
				minDist = dist
				jb1 = jb
				jb2 = otherJb
			}
		}
	}

	return jb1, jb2, minDist
}

func processLines(jbs []*JunctionBox) int {
	totalDist := 0

	allJBDists := findAllDistances(jbs)

	// Retrieve 10 shortest distances
	sortedDists := []int{}
	for dist := range allJBDists {
		sortedDists = append(sortedDists, dist)
	}
	fmt.Printf(" distances are: %v\n", sortedDists)

	sort.Ints(sortedDists)
	fmt.Printf("sorted shorted distances are: %v\n", sortedDists[0:10])

	// Do the 10 shortest connections
	for i := 0; i < 10; i++ {
		// find connections of closest isolated junction boxes (no connected together)
		jbDist := allJBDists[sortedDists[i]]
		jb1 := jbDist[0]
		jb2 := jbDist[1]
		if areInSameCircuit(jb1, jb2) {
			fmt.Printf("Junction boxes %s and %s are already in the same circuit, skipping\n", jb1, jb2)
			continue
		}

		connectTogether(jb1, jb2)
		fmt.Printf("Connecting junction boxes %s and %s (distance: %d)\n", jb1, jb2, sortedDists[i])
	}

	for i := -1; i < _circuitid; i++ {
		fmt.Printf("Circuit %d:\n", i)
		for _, jb := range jbs {
			if jb.CircuitID == i {
				fmt.Printf("-  %s\n", jb)
			}
		}
	}

	return totalDist
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
