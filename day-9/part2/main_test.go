package main

import (
	"testing"
)

var textGrid1 = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

var textGrid2 = `7,1
11,1
11,7
5,7
5,5
2,5
2,3
7,3`

var textGrid3 = `2,1
9,1
7,5
11,5
9,3
11,3
2,7
7,7`

var textGrid4 = `2,11
7,11
2,5
5,5
2,7
5,7
2,2
7,2`

func TestProcessLines(t *testing.T) {
	var positions []Position
	var want int
	var got int
	/**/
	test = true
	want = 24
	positions = parseLines(textGrid1)
	got = processLines(positions)
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
	/*
	   want = 35
	   positions = parseLines(textGrid2)
	   got = processLines(positions)

	   	if got != want {
	   		t.Errorf("processLines() = %v, want %v", got, want)
	   	}

	   want = 41
	   positions = parseLines(textGrid3)
	   got = processLines(positions)

	   	if got != want {
	   		t.Errorf("processLines() = %v, want %v", got, want)
	   	}

	   want = 41
	   positions = parseLines(textGrid4)
	   got = processLines(positions)

	   	if got != want {
	   		t.Errorf("processLines() = %v, want %v", got, want)
	   	}
	*/
}

func TestIsInShape(t *testing.T) {
	positions := parseLines(textGrid1)
	test = true
	displayBoard(positions)

	rowsRanges := computeRowsRanges(positions)

	tests := []struct {
		p    Position
		want bool
	}{
		{Position{X: 4, Y: 6}, false},
		{Position{X: 7, Y: 7}, false},
		{Position{X: 0, Y: 0}, false},
		{Position{X: 4, Y: 2}, false},
		{Position{X: 4, Y: 3}, true},
		{Position{X: 4, Y: 4}, true},
		{Position{X: 4, Y: 5}, true},
		{Position{X: 10, Y: 7}, true},
		{Position{X: 3, Y: 4}, true},
		{Position{X: 8, Y: 4}, true},
		{Position{X: 7, Y: 1}, true},
		{Position{X: 11, Y: 1}, true},
		{Position{X: 9, Y: 7}, true},
	}

	for _, tt := range tests {
		got := isInTheShape(tt.p, rowsRanges)
		if got != tt.want {
			t.Errorf("isInTheShape(%v) = %v, want %v", tt.p, got, tt.want)
		}
	}
}
