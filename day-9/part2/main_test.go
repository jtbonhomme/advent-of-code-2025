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

func TestProcessLines(t *testing.T) {
	var lines []Position
	var want int
	var got int
	/**/
	test = true
	want = 21
	lines = parseLines(textGrid1)
	got = processLines(lines)
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}

	want = 35
	lines = parseLines(textGrid2)
	got = processLines(lines)

	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}

	want = 41
	lines = parseLines(textGrid3)
	got = processLines(lines)

	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
	/**/
}
