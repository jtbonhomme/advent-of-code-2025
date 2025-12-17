package main

import (
	"testing"

	"github.com/jtbonhomme/advent-of-code-2025/utils"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Machine
	}{
		{
			name: "Test 0",
			args: args{
				line: "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			},
			want: Machine{lightDiagram: []bool{false, true, true, false}, buttonsWiring: [][]int{
				{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseLine(tt.args.line)
			for i, gotButtonWiring := range got.buttonsWiring {
				if !utils.EqualSlices(gotButtonWiring, tt.want.buttonsWiring[i]) {
					t.Errorf("%s buttonsWiring parseLine(%v) = %v, want %v", tt.name, tt.args.line, got, tt.want)
				}
			}
			if !utils.EqualSlices(got.lightDiagram, tt.want.lightDiagram) {
				t.Errorf("%s lightDiagram parseLine(%v) = %v, want %v", tt.name, tt.args.line, got, tt.want)
			}
		})
	}
}

func TestMinPresses(t *testing.T) {
	type args struct {
		machine Machine
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 0",
			args: args{
				machine: Machine{lightDiagram: []bool{false, true, true, false}, buttonsWiring: [][]int{
					{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1},
				}},
			},
			want: 2,
		},
		{
			name: "Test 1",
			args: args{
				machine: Machine{lightDiagram: []bool{false, false, false, true, false}, buttonsWiring: [][]int{
					{0, 2, 3, 4}, {2, 3}, {0, 4}, {0, 1, 2}, {1, 2, 3, 4},
				}},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				machine: Machine{lightDiagram: []bool{false, true, true, true, false, true}, buttonsWiring: [][]int{
					{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2},
				}},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPresses(tt.args.machine); got != tt.want {
				t.Errorf("%s minPresses() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

var textGrid = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

func TestProcessLines(t *testing.T) {
	want := 7

	lines := parseLines(textGrid)
	got := processLines(lines)
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
