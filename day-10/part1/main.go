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

type Machine struct {
	lightDiagram  []bool
	buttonsWiring [][]int
}

// Parse a single line of input
// ex.: [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
// ignore last part between brackets {} for now
// returns a lightDiagram object ([]bool) and a slice of buttons wiring ([][]int)
func parseLine(line string) Machine {
	var lightDiagram []bool
	var buttonsWiring [][]int

	//split line into parts
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return Machine{lightDiagram, buttonsWiring}
	}

	//parse light diagram
	diagramStr := parts[0]
	// [.##.] -> []bool{false, true, true, false}
	// trim the brackets
	diagramStr = strings.TrimPrefix(diagramStr, "[")
	diagramStr = strings.TrimSuffix(diagramStr, "]")
	lightDiagram = make([]bool, len(diagramStr))
	for i, ch := range diagramStr {
		if ch == '#' {
			lightDiagram[i] = true
		} else {
			lightDiagram[i] = false
		}
	}

	//parse buttons wiring
	for _, part := range parts[1 : len(parts)-1] {
		// (x,y) -> []int{x, y}
		part = strings.TrimPrefix(part, "(")
		part = strings.TrimSuffix(part, ")")
		coords := strings.Split(part, ",")
		if len(coords) == 0 {
			continue
		}

		var buttonWire []int
		for _, coord := range coords {
			var val int
			fmt.Sscanf(coord, "%d", &val)
			buttonWire = append(buttonWire, val)
		}

		buttonsWiring = append(buttonsWiring, buttonWire)
	}

	return Machine{lightDiagram, buttonsWiring}
}

func parseLines(i string) []Machine {
	res := []Machine{}
	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		machine := parseLine(line)
		res = append(res, machine)
	}

	return res
}

type State struct {
	lights  []bool
	presses map[int]State
}

func minPresses(machine Machine) int {
	// Build a tree starting with no button pressed
	// then apply all combinations of button presses
	// until the light diagram matches the target or returns to initial state (no buttons pressed)

	statesTree := State{
		lights:  make([]bool, len(machine.lightDiagram)),
		presses: make(map[int]State), // key is the index of buttonWire slice
	}

	for i, buttonWire := range machine.buttonsWiring {
		// Apply button wire to initialize
		newState := make([]bool, len(machine.lightDiagram))
		for _, lightIndex := range buttonWire {
			newState[lightIndex] = !newState[lightIndex]
		}
		statesTree.presses[i] = State{
			lights:  newState,
			presses: make(map[int]State),
		}
	}

	return len(machine.buttonsWiring)
}

func processLines(machines []Machine) int {
	var result int

	for _, machine := range machines {
		// For now, just count the number of buttons
		minPresses := minPresses(machine)
		result += minPresses
	}

	return result
}

func run(i string) int {
	machines := parseLines(i)

	answer := processLines(machines)

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
