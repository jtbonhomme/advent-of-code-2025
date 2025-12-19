package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"

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
	lightDiagram  LightDiagram
	buttonsWiring [][]int
}

// Parse a single line of input
// ex.: [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
// ignore last part between brackets {} for now
// returns a lightDiagram object (lightdiagram) and a slice of buttons wiring ([][]int)
func parseLine(line string) Machine {
	var lightDiagram LightDiagram
	var buttonsWiring [][]int

	//split line into parts
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return Machine{lightDiagram, buttonsWiring}
	}

	//parse light diagram
	diagramStr := parts[0]
	// [.##.] -> LightDiagram{false, true, true, false}
	// trim the brackets
	diagramStr = strings.TrimPrefix(diagramStr, "[")
	diagramStr = strings.TrimSuffix(diagramStr, "]")
	lightDiagram = make(LightDiagram, len(diagramStr))
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

func applyButtonPress(currentLights LightDiagram, buttonWire []int) LightDiagram {
	newLights := make(LightDiagram, len(currentLights))
	copy(newLights, currentLights)

	for _, lightIndex := range buttonWire {
		newLights[lightIndex] = !newLights[lightIndex]
	}

	return newLights
}

func buildStateTree(lights LightDiagram, machine Machine, buttonWiringIndexes []int, steps int) (int, bool) {
	if lights.IsNull() && steps > 0 {
		return steps, false
	}

	if steps > len(machine.buttonsWiring)*len(machine.buttonsWiring) {
		return -1, false
	}

	var ok bool
	var totalSteps int = -1
	for i, buttonWire := range machine.buttonsWiring {
		if slices.Contains(buttonWiringIndexes, i) {
			continue
		}
		// Apply button wire to initialize
		newLights := applyButtonPress(lights, buttonWire)
		if utils.EqualSlices(lights, machine.lightDiagram) {
			return steps, true
		}

		totalSteps, ok = buildStateTree(newLights, machine, append(buttonWiringIndexes, i), steps+1)
		if ok && totalSteps < minStepsGlobal {
			minStepsGlobal = totalSteps
		}
	}

	return totalSteps, ok
}

var minStepsGlobal int = -1

func minPresses(machine Machine) int {
	// Build a tree starting with no button pressed
	// then apply all combinations of button presses
	// until the light diagram matches the target or returns to initial state (no buttons pressed)
	minStepsGlobal = len(machine.lightDiagram) * len(machine.buttonsWiring)
	_, _ = buildStateTree(make(LightDiagram, len(machine.lightDiagram)), machine, []int{}, 0)
	return minStepsGlobal
}

func processLines(machines []Machine) int {
	var result int

	for i, machine := range machines {
		// For now, just count the number of buttons
		fmt.Printf("[%d] min presses = ", i)
		minPresses := minPresses(machine)
		fmt.Printf("%d\n", minPresses)
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
