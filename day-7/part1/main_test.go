package main

import (
	"testing"

	"github.com/jtbonhomme/advent-of-code-2025/utils"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line  string
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int
		wantSplit  int
	}{
		{
			name: "Test 0",
			args: args{
				line:  ".......S.......",
				input: []int{},
			},
			wantOutput: []int{7},
			wantSplit:  0,
		},
		{
			name: "Test 1",
			args: args{
				line:  ".......^.......",
				input: []int{7},
			},
			wantOutput: []int{6, 8},
			wantSplit:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, gotSplit := parseLine(tt.args.line, tt.args.input)
			if !utils.EqualSlices(gotOutput, tt.wantOutput) {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, gotOutput, tt.wantOutput)
			}
			if gotSplit != tt.wantSplit {
				t.Errorf("%s parseLine(%v) split = %v, want %v", tt.name, tt.args.line, gotSplit, tt.wantSplit)
			}
		})
	}
}

var textGrid = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func TestProcessLines(t *testing.T) {
	want := 21

	lines, totalSplit := parseLines(textGrid)
	got := processLines(lines, totalSplit)
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
