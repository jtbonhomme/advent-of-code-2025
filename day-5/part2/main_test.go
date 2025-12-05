package main

import (
	"testing"
)

func TestParseRangeLine(t *testing.T) {
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
				line: "1-3",
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseRangeLine(tt.args.line)
			if !equalSlices(got, tt.want) {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, got, tt.want)
			}
		})
	}
}

var textGrid = `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestParseLines(t *testing.T) {
	wantRanges := [][]int{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}

	gotFreshIDs := parseLines(textGrid)
	for i, r := range gotFreshIDs {
		if !equalSlices(r, wantRanges[i]) {
			t.Errorf("parseLines() got range %v want %v", r, wantRanges[i])
		}
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

func TestAdd(t *testing.T) {
	// Implement test for IDRanges.add method
	idRanges := []IDRange{}
	idRanges = addRange(idRanges, IDRange{start: 1, end: 5})
	idRanges = addRange(idRanges, IDRange{start: 6, end: 10})
	idRanges = addRange(idRanges, IDRange{start: 4, end: 7})   // overlaps with first two
	idRanges = addRange(idRanges, IDRange{start: 12, end: 15}) // overlaps with first two

	want := []IDRange{
		{start: 1, end: 10},
		{start: 12, end: 15},
	}

	if len(idRanges) != len(want) {
		t.Fatalf("Expected %d ranges, got %d", len(want), len(idRanges))
	}

	//for i, r := range idRanges.ranges {
	//	if r.start != want[i].start || r.end != want[i].end {
	//		t.Errorf("Range %d: got (%d, %d), want (%d, %d)", i, r.start, r.end, want[i].start, want[i].end)
	//	}
	//}
}

func TestProcessLines(t *testing.T) {
	want := 14

	freshIDs := parseLines(textGrid)
	got := processLines(freshIDs)
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
