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

func computeNewDial(currentDial int, i int, dist int) (int, int) {
	zero := 0

	if dist > 99 {
		zero += dist / 100
		dist = dist % 100
	}

	newDial := currentDial + (i * dist)

	if newDial > 99 {
		newDial -= 100
		zero++
	}
	if newDial < 0 && currentDial != 0 {
		zero++
	}
	if newDial < 0 {
		newDial += 100
	}
	if newDial == 0 && i == -1 {
		zero++
	}

	return newDial, zero
}

func run(i string) int {
	var answer, zero int
	var dial int = 50

	fmt.Println("dial: ", dial)
	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		i, dist := parseLine(line)

		dial, zero = computeNewDial(dial, i, dist)

		if dial < 0 || dial > 99 {
			panic("dial out of bounds")
		}
		fmt.Println("dial:", dial)
		fmt.Println("zero:", zero)
		if dial == 0 {
			answer += zero
		}
	}

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
