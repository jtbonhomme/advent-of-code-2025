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

var textGrid2 = `2,1
9,1
7,5
11,5
9,3
11,3
2,7
7,7`

func TestProcessLines(t *testing.T) {
	want1 := 150
	lines1 := parseLines(textGrid1)
	got1 := processLines(lines1)
	if got1 != want1 {
		t.Errorf("processLines() = %v, want %v", got1, want1)
	}

	want2 := 150
	lines2 := parseLines(textGrid2)
	got2 := processLines(lines2)
	if got2 != want2 {
		t.Errorf("processLines() = %v, want %v", got2, want2)
	}
}
