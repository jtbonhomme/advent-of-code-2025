package main

import (
	"fmt"
	"testing"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{
			name: "Test 0",
			args: args{
				line: "906,360,560",
			},
			want: Position{X: 906, Y: 360, Z: 560},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseLine(tt.args.line)
			if got != tt.want {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, got, tt.want)
			}
		})
	}
}

var testParse string = `162,817,812
57,618,57
906,360,560`

func TestParseLines(t *testing.T) {
	want := []*JunctionBox{
		{Position: Position{X: 162, Y: 817, Z: 812}},
		{Position: Position{X: 57, Y: 618, Z: 57}},
		{Position: Position{X: 906, Y: 360, Z: 560}},
	}

	got := parseLines(testParse)
	for i, jb := range want {
		if equalPosition(got[i].Position, jb.Position) == false {
			t.Errorf("parseLines() line %d = %v, want %v", i, got[i], want[i])
		}
	}
}

var textGrid = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
431,825,988
542,29,236
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

var smallTextGrid = `162,817,812
431,825,988
984,92,344
425,690,689`

func TestFindClosestNonConnectedBoxes(t *testing.T) {
	lines := parseLines(smallTextGrid)
	jb1, jb2, got := findClosestNonConnectedBoxes(lines)
	if got != 100427 {
		t.Errorf("findClosestNonConnectedBoxes() = %v, want %v", got, 100427)
	}
	connectTogether(jb1, jb2)

	jb3, jb4, _ := findClosestNonConnectedBoxes(lines)
	connectTogether(jb3, jb4)

	_, _, _ = findClosestNonConnectedBoxes(lines)

}

/*
// jb1 -> jb2 -> ... -> jbN -> jb1
func TestIsClosedCircuit(t *testing.T) {
	lines := parseLines(smallTextGrid)
	circuit := []*JunctionBox{}

	jb1, jb2, _ := findClosestNonConnectedBoxes(lines)
	if isClosedCircuit(circuit) == true {
		t.Errorf("isClosedCircuit() = true, want false")
	}

	connectTogether(jb1, jb2)
	if isClosedCircuit(circuit) == true {
		t.Errorf("isClosedCircuit() = ftrue, want false")
	}

	addToCircuit(&circuit, jb1)
	addToCircuit(&circuit, jb2)
	if isClosedCircuit(circuit) == true {
		t.Errorf("isClosedCircuit() = true, want false")
	}

	jb3, jb4, _ := findClosestNonConnectedBoxes(lines)
	if isClosedCircuit(circuit) == true {
		t.Errorf("isClosedCircuit() = true, want false")
	}
	connectTogether(jb3, jb4)
	if isClosedCircuit(circuit) == true {
		t.Errorf("isClosedCircuit() = true, want false")
	}

	addToCircuit(&circuit, jb3)
	addToCircuit(&circuit, jb4)
	if isClosedCircuit(circuit) == true {
		t.Errorf("isClosedCircuit() = true, want false")
	}

	connectTogether(jb4, jb1)
	if isClosedCircuit(circuit) == false {
		t.Errorf("isClosedCircuit() = false, want true")
	}
}*/

func TestAreConnected(t *testing.T) {
	lines := parseLines(textGrid)

	if areConnected(lines[0], lines[1]) == true {
		t.Errorf("areConnected() = false, want true")
	}

	connectTogether(lines[0], lines[1])

	if areConnected(lines[0], lines[1]) == false {
		t.Errorf("areConnected() = false, want true")
	}

	if areConnected(lines[2], lines[1]) == true {
		t.Errorf("areConnected() = true, want false")
	}

	if areConnected(lines[2], lines[0]) == true {
		t.Errorf("areConnected() = true, want false")
	}

	connectTogether(lines[2], lines[1])

	if areConnected(lines[2], lines[1]) == false {
		t.Errorf("areConnected() = false, want true")
	}

	if areConnected(lines[2], lines[0]) == false {
		t.Errorf("areConnected() = false, want true")
	}
}

func TestProcessLines(t *testing.T) {
	lines := parseLines(textGrid)
	totalDist := processLines(lines)
	fmt.Printf("totalDist=%d\n", totalDist)
}
