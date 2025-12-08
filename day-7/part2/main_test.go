package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line  string
		input map[int]int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput map[int]int
	}{
		{
			name: "Test 0",
			args: args{
				line:  ".......S.......",
				input: map[int]int{},
			},
			wantOutput: map[int]int{7: 1},
		},
		{
			name: "Test 0.5",
			args: args{
				line:  "...............",
				input: map[int]int{7: 1},
			},
			wantOutput: map[int]int{7: 1},
		},
		{
			name: "Test 1",
			args: args{
				line:  ".......^.......",
				input: map[int]int{7: 1},
			},
			wantOutput: map[int]int{6: 1, 8: 1},
		},
		{
			name: "Test 2",
			args: args{
				line:  "......^.^......",
				input: map[int]int{6: 1, 8: 1},
			},
			wantOutput: map[int]int{5: 1, 7: 2, 9: 1},
		},
		{
			name: "Test 3",
			args: args{
				line:  ".....^.^.^.....",
				input: map[int]int{5: 1, 7: 2, 9: 1},
			},
			wantOutput: map[int]int{4: 1, 6: 3, 8: 3, 10: 1},
		},
		{
			name: "Test 4",
			args: args{
				line:  "....^.^...^....",
				input: map[int]int{4: 1, 6: 3, 8: 3, 10: 1},
			},
			wantOutput: map[int]int{3: 1, 5: 4, 7: 3, 8: 3, 9: 1, 11: 1},
		},
		{
			name: "Test 5",
			args: args{
				line:  "...^.^...^.^...",
				input: map[int]int{3: 1, 5: 4, 7: 3, 8: 3, 9: 1, 11: 1},
			},
			wantOutput: map[int]int{2: 1, 4: 5, 6: 4, 7: 3, 8: 4, 10: 2, 12: 1},
		},
		{
			name: "Test 6",
			args: args{
				line:  "..^...^.....^..",
				input: map[int]int{2: 1, 4: 5, 6: 4, 7: 3, 8: 4, 10: 2, 12: 1},
			},
			wantOutput: map[int]int{1: 1, 3: 1, 4: 5, 5: 4, 7: 7, 8: 4, 10: 2, 11: 1, 13: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput := parseLine(tt.args.line, tt.args.input)
			result := reflect.DeepEqual(gotOutput, tt.wantOutput)
			if !result {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, gotOutput, tt.wantOutput)
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
	want := 40

	lines := parseLines(textGrid)
	got := processLines(lines)
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
