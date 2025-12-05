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

func TestParseIDLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 0",
			args: args{
				line: "34",
			},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseProductIDLine(tt.args.line)
			if got != tt.want {
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
	wantIDs := []int{
		1, 5, 8, 11, 17, 32,
	}
	wantRanges := [][]int{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}

	gotFreshIDs, gotProductIDs := parseLines(textGrid)
	if !equalSlices(gotProductIDs, wantIDs) {
		t.Errorf("parseLines() got ids %v want %v", gotProductIDs, wantIDs)
	}
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
