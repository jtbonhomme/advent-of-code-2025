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
		want []int
	}{
		{
			name: "Test 0",
			args: args{
				line: "..@@.@@@@.",
			},
			want: []int{0, 0, 1, 1, 0, 1, 1, 1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseLine(tt.args.line)
			if !equalSlices(got, tt.want) {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, got, tt.want)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var textGrid = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

// Example grid:
//    0. 1. 2. 3. 4. 5. 6. 7. 8. 9
//  +-----------------------------
// 0| 0  0  1  1  0  1  1  1  1  0
// 0| 1  1  1  0  1  0  1  0  1  1
// 0| 1  1  1  1  1  0  1  0  1  1
// 0| 1  0  1  1  1  1  0  0  1  0
// 0| 1  1  0  1  1  1  1  0  1  1
// 0| 0  1  1  1  1  1  1  1  0  1
// 0| 0  1  0  1  0  1  0  1  1  1
// 0| 1  0  1  1  1  0  1  1  1  1
// 0| 0  1  1  1  1  1  1  1  1  0
// 0| 1  0  1  0  1  1  1  0  1  0

func TestParseLines(t *testing.T) {
	want := [][]int{
		{0, 0, 1, 1, 0, 1, 1, 1, 1, 0},
		{1, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 0, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 1, 0, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0, 1, 0},
	}

	got := parseLines(textGrid)
	for i := range want {
		if !equalSlices(got[i], want[i]) {
			t.Errorf("parseLines() line %d = %v, want %v", i, got[i], want[i])
		}
	}
}

func TestSumAdjacentCells(t *testing.T) {
	grid := parseLines(textGrid)

	type args struct {
		x, y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 0,0",
			args: args{x: 0, y: 0},
			want: 2,
		},
		{
			name: "Test 1,0",
			args: args{x: 1, y: 0},
			want: 4,
		},
		{
			name: "Test 2,0",
			args: args{x: 2, y: 0},
			want: 3,
		},
		{
			name: "Test 3,0",
			args: args{x: 3, y: 0},
			want: 3,
		},
		{
			name: "Test 7,0",
			args: args{x: 7, y: 0},
			want: 4,
		},
		{
			name: "Test 8,0",
			args: args{x: 8, y: 0},
			want: 3,
		},
		{
			name: "Test 0,1",
			args: args{x: 0, y: 1},
			want: 3,
		},
		{
			name: "Test 1,1",
			args: args{x: 1, y: 1},
			want: 6,
		},
		{
			name: "Test 6,1",
			args: args{x: 6, y: 1},
			want: 4,
		},
		{
			name: "Test 6,2",
			args: args{x: 6, y: 2},
			want: 2,
		},
		{
			name: "Test 5,3",
			args: args{x: 5, y: 3},
			want: 6,
		},
		{
			name: "Test 0,9",
			args: args{x: 0, y: 9},
			want: 1,
		},
		{
			name: "Test 4,9",
			args: args{x: 4, y: 9},
			want: 4,
		},
		{
			name: "Test 9,9",
			args: args{x: 9, y: 9},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumAdjacentCells(grid, tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("%s sumAdjacentCells(%d, %d) = %d, want %d", tt.name, tt.args.x, tt.args.y, got, tt.want)
			}
		})
	}
}

/*
..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.

..XX.XX@X.
X@@.@.@.@@
@@@@@.X.@@
@.@@@@..@.
X@.@@@@.@X
.@@@@@@@.@
.@.@.@.@@@
X.@@@.@@@@
.@@@@@@@@.
X.X.@@@.X.
*/
func TestGetPaperRolls(t *testing.T) {
	grid := parseLines(textGrid)
	for y, line := range grid {
		for x, cell := range line {
			// no roll paper at this position
			if cell == 0 {
				fmt.Printf(".")
				continue
			}

			n := sumAdjacentCells(grid, x, y)
			if n < 4 {
				fmt.Printf("X")
			} else {
				fmt.Printf("@")
			}
		}
		fmt.Printf("\n")
	}
}
