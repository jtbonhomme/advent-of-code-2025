package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"time"

	//    "regexp"
	//    "strconv"
	"strings"

	"github.com/jtbonhomme/advent-of-code-2025/utils"
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

func (machine *Machine) buildStateTree(currentJoltage Joltage, steps int) (int, bool) {
	if steps > len(machine.buttonsWiring)*len(machine.buttonsWiring) {
		return -1, false
	}

	if currentJoltage.IsGreater(machine.joltage) {
		return -1, false
	}

	var ok bool
	var totalSteps int = -1
	for _, buttonWire := range machine.buttonsWiring {
		// Apply button wire to initialize
		newJoltage := applyButtonPress(currentJoltage, buttonWire)
		if utils.EqualSlices(newJoltage, machine.joltage) {
			return steps, true
		}

		totalSteps, ok = machine.buildStateTree(newJoltage, steps+1)
		if ok && totalSteps < machine.minPresses {
			machine.minPresses = totalSteps
		}
		if !ok {
			return -1, false
		}
	}

	return totalSteps, ok
}

func minPresses(machine Machine) int {
	// Build a tree starting with no button pressed
	// then apply all combinations of button presses
	// until the light diagram matches the target or returns to initial state (no buttons pressed)
	_, _ = machine.buildStateTree(make(Joltage, len(machine.joltage)), 0)
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
