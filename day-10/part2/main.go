package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"time"

	//    "regexp"
	//    "strconv"
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

type LightDiagram []bool
type Joltage []int

func (ld LightDiagram) IsNull() bool {
	for _, light := range ld {
		if light {
			return false
		}
	}

	return true
}

func (ld LightDiagram) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	for _, light := range ld {
		if light {
			sb.WriteString("#")
		} else {
			sb.WriteString(".")
		}
	}

	sb.WriteString("]")

	return sb.String()
}

type Machine struct {
	buttonsWiring [][]int
	joltage       Joltage
	minPresses    int
}

// Parse a single line of input
// ex.: [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
// ignore last part between brackets {} for now
// returns a lightDiagram object (lightdiagram) and a slice of buttons wiring ([][]int)
func parseLine(line string) Machine {
	var joltage Joltage
	var buttonsWiring [][]int

	//split line into parts
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return Machine{buttonsWiring, joltage, -1}
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

	//parse joltage
	joltagePart := parts[len(parts)-1]
	joltagePart = strings.TrimPrefix(joltagePart, "{")
	joltagePart = strings.TrimSuffix(joltagePart, "}")
	joltageStrs := strings.Split(joltagePart, ",")
	for _, js := range joltageStrs {
		var val int
		fmt.Sscanf(js, "%d", &val)
		joltage = append(joltage, val)
	}

	return Machine{buttonsWiring, joltage, len(buttonsWiring) * len(buttonsWiring)}
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

func applyButtonPress(currentJoltage Joltage, buttonWire []int) Joltage {
	newJoltage := make(Joltage, len(currentJoltage))
	copy(newJoltage, currentJoltage)

	for _, lightIndex := range buttonWire {
		newJoltage[lightIndex]++
	}

	return newJoltage
}

func (j Joltage) IsGreater(other Joltage) bool {
	if len(j) != len(other) {
		return false
	}

	for i := range j {
		if j[i] > other[i] {
			return true
		}
	}

	return false
}

func (j Joltage) Equals(other Joltage) bool {
	if len(j) != len(other) {
		return false
	}

	for i := range j {
		if j[i] != other[i] {
			return false
		}
	}

	return true
}

func displayAction(pressedButtons []int, newJoltage Joltage, machine Machine) {
	fmt.Printf("Position.")
	for pos := range machine.joltage {
		fmt.Printf(" %d ", pos)
	}
	fmt.Println()

	fmt.Printf("Joltage..")
	for _, j := range machine.joltage {
		fmt.Printf(" %d ", j)
	}
	fmt.Println()

	fmt.Printf("Try......")
	for _, j := range newJoltage {
		fmt.Printf(" %d ", j)
	}

	actions := []string{}
	for i, b := range pressedButtons {
		actions = append(actions, fmt.Sprintf("b%d x %d", i, b))
	}
	fmt.Println(strings.Join(actions, " + "))
}

// Algorithm to find the minimum button presses to reach the target joltage
// First rank by descending order the positions from the highest joltage needed to the lowest.
// For example if the target joltage is {3,5,4,7}, the order of positions to satisfy is 3,1,2,0
// Then start a successive rounds to find all presses combinatory that satisfy position 3 (without overpassing other position's joltage needs), then position 1, then position 2, then position 0.
// For each position, identify the buttons that can contribute to increase the joltage at that position without overpassing other positions' needs.
// Build a tree of all combinations of button presses that satisfy the current position, then for each leaf node, continue to the next position until all positions are satisfied.
// Keep track of the minimum number of presses found.
func (machine *Machine) getMinPresses() {
	actions := make([]int, len(machine.buttonsWiring))
	currentJoltage := make(Joltage, len(machine.joltage))
	newJoltage := make(Joltage, len(currentJoltage))
	copy(newJoltage, currentJoltage)

	actions[5] = 2 // press button 5 two times
	actions[4] = 1 // press button 4 one time
	actions[3] = 3 // press button 3 three times
	actions[1] = 3 // press button 1 three times
	actions[0] = 1 // press button 0 one time
	steps := 0
	for i, b := range actions {
		for j := 0; j < b; j++ {
			newJoltage = applyButtonPress(newJoltage, machine.buttonsWiring[i])
			steps++
		}
	}

	displayAction(actions, newJoltage, *machine)
	if newJoltage.Equals(machine.joltage) {
		fmt.Printf("Solved in %d steps!\n", steps)
		machine.minPresses = steps
	} else {
		fmt.Printf("Not solved.\n")
	}
}

func minPresses(machine Machine) int {
	// Build a tree starting with no button pressed
	// then apply all combinations of button presses
	// until the light diagram matches the target or returns to initial state (no buttons pressed)
	machine.getMinPresses()

	return machine.minPresses
}

func processLines(machines []Machine) int {
	var result int

	for i, machine := range machines {
		start := time.Now()
		// For now, just count the number of buttons
		fmt.Printf("[%d] min presses = ", i)
		minPresses := minPresses(machine)
		fmt.Printf("%d (took %v)\n", minPresses, time.Since(start))
		result += minPresses
	}

	return result
}

func run(i string) int {
	start := time.Now()
	fmt.Printf("Parsing input...\n")
	machines := parseLines(i)

	answer := processLines(machines)
	fmt.Printf("Completed in %v\n", time.Since(start))
	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
