package main

import (
	"testing"
)

var textGrid = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestProcessLines(t *testing.T) {
	want := 50

	lines := parseLines(textGrid)
	got := processLines(lines)
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
