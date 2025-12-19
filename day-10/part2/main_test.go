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
			want: Machine{buttonsWiring: [][]int{
				{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1},
			},
				joltage: []int{3, 5, 4, 7},
			},
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
			if !utils.EqualSlices(got.joltage, tt.want.joltage) {
				t.Errorf("%s joltage parseLine(%v) = %v, want %v", tt.name, tt.args.line, got, tt.want)
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
		/*
			// [..#.###.#.] (1,6,8) (1,2,3,4,6) (1,4,7,8) (1,2,3,4,5,6,9) (6,7,9) (1,2,3,5,6,8,9) (0,1,3,5,8,9) (0,1,2,4,5,8,9) (5,6,8,9) (0,9) (1,2) (2,4,5,7) (1,3,4,5,7,8,9) {29,286,259,227,83,256,242,61,260,263}
			{
				name: "Test 4",
				args: args{
					machine: Machine{
						joltage: []int{29, 286, 259, 227, 83, 256, 242, 61, 260, 263},
						buttonsWiring: [][]int{
							{1, 6, 8}, {1, 2, 3, 4, 6}, {1, 4, 7, 8}, {1, 2, 3, 4, 5, 6, 9}, {6, 7, 9},
							{1, 2, 3, 5, 6, 8, 9}, {0, 1, 3, 5, 8, 9}, {0, 1, 2, 4, 5, 8, 9}, {5, 6, 8, 9},
							{0, 9}, {1, 2}, {2, 4, 5, 7}, {1, 3, 4, 5, 7, 8, 9},
						},
						minPresses: 10000,
					}},
				want: 5,
			},
			// [#.##...] (1,3,4,5,6) (1,3) (0,5) (0,1,2,4,5,6) (2,3,5) {16,7,14,15,3,27,3}
			{
				name: "Test 3",
				args: args{
					machine: Machine{
						joltage: []int{16, 7, 14, 15, 3, 27, 3},
						buttonsWiring: [][]int{
							{1, 3, 4, 5, 6}, {1, 3}, {0, 5}, {0, 1, 2, 4, 5, 6}, {2, 3, 5},
						},
						minPresses: 10000,
					}},
				want: 2,
			},

			// [.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
			{
				name: "Test 2",
				args: args{
					machine: Machine{joltage: []int{10, 11, 11, 5, 10, 5}, buttonsWiring: [][]int{
						{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2},
					},
						minPresses: 10000,
					}},
				want: 2,
			},
			// [...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
			{
				name: "Test 1",
				args: args{
					machine: Machine{joltage: []int{7, 5, 12, 7, 2}, buttonsWiring: [][]int{
						{0, 2, 3, 4}, {2, 3}, {0, 4}, {0, 1, 2}, {1, 2, 3, 4},
					},
						minPresses: 10000,
					}},
				want: 3,
			},
		*/
		// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
		{
			name: "Test 0",
			args: args{
				machine: Machine{joltage: []int{3, 5, 4, 7}, buttonsWiring: [][]int{
					{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1},
				},
					minPresses: 10000,
				}},
			want: 2,
		},
	}

	test = true
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
	test = true

	lines := parseLines(textGrid)
	got := processLines(lines)
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
