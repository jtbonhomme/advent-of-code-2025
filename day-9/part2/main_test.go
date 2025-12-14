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
	colsRanges = computeColsRanges(positions)

	tests := []struct {
		p    Position
		want bool
	}{
		{Position{X: 3, Y: 4}, true},
		{Position{X: 4, Y: 4}, true},
		{Position{X: 8, Y: 4}, true},
		{Position{X: 3, Y: 2}, false},
		{Position{X: 4, Y: 2}, false},
		{Position{X: 8, Y: 2}, true},
		{Position{X: 2, Y: 7}, false},
		{Position{X: 9, Y: 3}, true},
		{Position{X: 7, Y: 7}, false},
		{Position{X: 0, Y: 1}, false},
		{Position{X: 0, Y: 7}, false},
		{Position{X: 6, Y: 7}, false},
		{Position{X: 4, Y: 5}, true},
		{Position{X: 10, Y: 7}, true},
		{Position{X: 4, Y: 6}, false},
		{Position{X: 0, Y: 0}, false},
		{Position{X: 4, Y: 2}, false},
		{Position{X: 4, Y: 3}, true},
		{Position{X: 7, Y: 1}, true},
		{Position{X: 11, Y: 1}, true},
		{Position{X: 11, Y: 7}, true},
		{Position{X: 9, Y: 7}, true},
	}

	for _, tt := range tests {
		got := isInTheShape(tt.p, rowsRanges, colsRanges)
		if got != tt.want {
			t.Errorf("isInTheShape(%v) = %v, want %v", tt.p, got, tt.want)
		}
	}
}

func TestExtendRows(t *testing.T) {
	rowsRanges := map[int][]int{
		1: {7, 11},
		3: {2, 7},
		5: {2, 9},
		7: {9, 11},
	}

	colsRanges := map[int][]int{
		2:  {3, 5},
		7:  {1, 3},
		9:  {5, 7},
		11: {1, 7},
	}
	extendedRows := extendRows(rowsRanges, colsRanges)

	wantExtendedRows := map[int][]int{
		1: {7, 11},
		3: {2, 11},
		5: {2, 11},
		7: {9, 11},
	}

	for row, wantRange := range wantExtendedRows {
		gotRange, exists := extendedRows[row]
		if !exists {
			t.Errorf("extendRows() missing row %v, want range %v", row, wantRange)
			continue
		}
		if gotRange[0] != wantRange[0] || gotRange[1] != wantRange[1] {
			t.Errorf("extendRows() row %v = %v, want %v", row, gotRange, wantRange)
		}
	}
}

func TestExtendCols(t *testing.T) {
	rowsRanges := map[int][]int{
		1: {7, 11},
		3: {2, 7},
		5: {2, 9},
		7: {9, 11},
	}

	colsRanges := map[int][]int{
		2:  {3, 5},
		7:  {1, 3},
		9:  {5, 7},
		11: {1, 7},
	}
	extendedCols := extendCols(rowsRanges, colsRanges)
	wantExtendedCols := map[int][]int{
		2:  {3, 5},
		7:  {1, 5},
		9:  {1, 7},
		11: {1, 7},
	}

	for col, wantRange := range wantExtendedCols {
		gotRange, exists := extendedCols[col]
		if !exists {
			t.Errorf("extendCols() missing col %v, want range %v", col, wantRange)
			continue
		}
		if gotRange[0] != wantRange[0] || gotRange[1] != wantRange[1] {
			t.Errorf("extendCols() col %v = %v, want %v", col, gotRange, wantRange)
		}
	}
}
