package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"

	"strings"
)

//go:embed input.txt
var input string

func parseLine(line string) (int, int) {
	i := 1
	dir := line[0:1]
	fmt.Println("            line:", line)
	fmt.Println("            dir :", dir)
	if dir == "L" {
		i = -1
	}
	fmt.Println("            i   :", i)
	dist, _ := strconv.Atoi(line[1:])
	fmt.Println("            dist:", dist)

	return i, dist
}

func computeNewDial(currentDial int, i int, dist int) int {
	newDial := currentDial + (i * dist)

	if newDial > 99 {
		newDial -= 100
	}
	if newDial < 0 {
		newDial += 100
	}

	return newDial
}

func run(i string) int {
	var answer int
	var initialDial int = 50

	fmt.Println("dial: ", initialDial)
	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		i, dist := parseLine(line)

		initialDial = computeNewDial(initialDial, i, dist)

		fmt.Println("dial:", initialDial)
		if initialDial == 0 {
			answer++
		}
	}

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
